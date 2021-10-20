package handlers

import (
	"fmt"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/database"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/helpers"
	"time"
)

func getUnRecognizedArg() mdparser.WMarkDown {
	msg := mdparser.GetBold("Unrecognized arguments\n\n").AppendNormal("Supported args are\n")
	msg.AppendMonoThis("bans").AppendItalicThis(" - toggle SpamProtection Bans\n")
	msg.AppendMonoThis("detection").AppendItalicThis(" - toggle spam detection")
	return msg
}

func settingsHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	message := ctx.EffectiveMessage
	if chat.Type == PrivateChat {
		_, err := message.Reply(b, mdparser.GetItalic("Command only for group").ToString(), &gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
		return helpers.SendError(err, ctx, b)
	}
	data, err := database.GetChat(chat.Id)
	if err != nil {
		return helpers.SendError(err, ctx, b)
	}
	txt := mdparser.GetBold(chat.Title).AppendMono(fmt.Sprintf(" (%d)", chat.Id)).AppendNormal("\n\n")
	if data == nil {
		txt.AppendItalicThis("No data found")
	} else {
		txt.AppendBoldThis("SpamProtection Bans: ")
		if data.DoesAutoBan() {
			txt.AppendMonoThis("On")
		} else {
			txt.AppendMonoThis("Off")
		}
		txt.AppendNormalThis("\n")
		txt.AppendBoldThis("Spam Prediction: ")
		if data.DetectSpam() {
			txt.AppendMonoThis("On")
		} else {
			txt.AppendMonoThis("Off")
		}
		txt.AppendNormalThis("\n")
		txt.AppendNormalThis("\n").AppendBoldThis("Fetched on ").AppendBoldThis(time.Now().UTC().Format(TimeLayout))
	}
	_, err = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
	return helpers.SendError(err, ctx, b)
}

func toggleHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	chat := ctx.EffectiveChat
	args := ctx.Args()
	if chat.Type == PrivateChat {
		_, err := message.Reply(b, mdparser.GetItalic("Command only for group").ToString(), &gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
		return helpers.SendError(err, ctx, b)
	}
	data, err := database.GetChat(chat.Id)
	if err != nil {
		return helpers.SendError(err, ctx, b)
	}
	if data == nil {
		_, err := message.Reply(b, mdparser.GetItalic("No data found").ToString(), &gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
		return helpers.SendError(err, ctx, b)
	}
	if len(args) < 2 {
		_, err := message.Reply(b, getUnRecognizedArg().ToString(), &gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
		return helpers.SendError(err, ctx, b)
	}

	arg := args[1]
	switch arg {
	case "bans":
		database.InsertChat(chat.Id, !data.DoesAutoBan(), data.DetectSpam(), data.SpamAction)
		if !data.DoesAutoBan() {
			_, err := message.Reply(b, mdparser.GetNormal("SpamProtection Bans enabled in ").AppendBold(chat.Title).ToString(),
				&gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
			return helpers.SendError(err, ctx, b)
		} else {
			_, err := message.Reply(b, mdparser.GetNormal("SpamProtection Bans disabled in ").AppendBold(chat.Title).ToString(),
				&gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
			return helpers.SendError(err, ctx, b)
		}
	case "detection":
		database.InsertChat(chat.Id, data.DoesAutoBan(), !data.DetectSpam(), data.SpamAction)
		if !data.DetectSpam() {
			_, err := message.Reply(b, mdparser.GetNormal("Spam detection enabled in ").AppendBold(chat.Title).ToString(),
				&gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
			return helpers.SendError(err, ctx, b)
		} else {
			_, err := message.Reply(b, mdparser.GetNormal("Spam detection disabled in ").AppendBold(chat.Title).ToString(),
				&gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
			return helpers.SendError(err, ctx, b)
		}
	default:
		_, err := message.Reply(b, getUnRecognizedArg().ToString(),
			&gotgbot.SendMessageOpts{ParseMode: MarkdownV2})
		return helpers.SendError(err, ctx, b)
	}
}
