package helpers

import (
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/core"
	"strings"
)

func SendError(err error, ctx *ext.Context, b *gotgbot.Bot) error {
	if err == nil {
		return nil
	}
	msg := mdparser.GetBold("Error: ").AppendItalic(err.Error())
	txt := msg.ToString()
	m := strings.ReplaceAll(txt, b.Token, "token")
	m = strings.ReplaceAll(m, core.Data.CoffeeHouseKey, "cf_key")
	_, er := b.SendMessage(ctx.EffectiveChat.Id, m, &gotgbot.SendMessageOpts{ParseMode: "markdownv2"})
	if er != nil {
		core.SUGARED.Error(err)
	}
	return err
}
