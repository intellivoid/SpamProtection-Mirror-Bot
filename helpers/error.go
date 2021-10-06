package helpers

import (
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/core"
	"strings"
)

func SendError(err error, ctx *ext.Context, b *gotgbot.Bot) {
	msg := mdparser.GetBold("Error: ").AppendItalic(err.Error())
	_, er := b.SendMessage(ctx.EffectiveChat.Id, strings.ReplaceAll(msg.ToString(), b.Token, "token"), &gotgbot.SendMessageOpts{ParseMode: "markdownv2"})
	if er != nil {
		core.SUGARED.Error(err)
	}
}
