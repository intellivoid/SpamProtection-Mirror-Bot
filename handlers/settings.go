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
