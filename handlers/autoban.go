package handlers

import (
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/Intellivoid/Intellivoid.SpamProtection-go/spamProtection"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/helpers"
)

func autoBanHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	user := ctx.EffectiveUser
	data, err := spamProtection.GetInfoByID(user.Id)
	if err != nil {
		helpers.SendError(err, ctx, b)
		return err
	}
	if !data.Success {
		return nil
	}
	if data.IsBlacklisted() {
		_, err := b.BanChatMember(chat.Id, user.Id, nil)
		if err != nil {
			// helpers.SendError(err, ctx, b)
			return err
		}
		txt := mdparser.GetUserMention(user.FirstName, user.Id).AppendBold(" was blacklisted and has been banned!").AppendNormal("\n")
		txt = txt.AppendBold("Reason").AppendNormal(": ").AppendMono(data.Results.Attributes.BlacklistReason)
		_, err = b.SendMessage(chat.Id, txt.ToString(), &gotgbot.SendMessageOpts{ParseMode: "markdownv2"})
		return err
	}
	return nil
}

func msgFilter(msg *gotgbot.Message) bool {
	return msg.From != nil
}