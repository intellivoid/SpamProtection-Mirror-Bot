package handlers

import (
	"fmt"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func startHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()
	m := ctx.EffectiveMessage
	if len(args) > 1 && args[1] == "help" {
		_, err := b.SendMessage(m.Chat.Id, getHelp(m.From), &gotgbot.SendMessageOpts{ParseMode: "markdownv2"})
		return err
	}
	txt := mdparser.GetBold("Hi, I'm " + b.FirstName).AppendNormal("\n\n")
	txt = txt.AppendNormal(`I'm a mirror of @SpamProtectionBot, this means`)
	txt = txt.AppendNormal(` I enforce all bans made through that bot and also provide user information related to it`)
	_, err := b.SendMessage(ctx.Message.Chat.Id, txt.ToString(), &gotgbot.SendMessageOpts{ParseMode: "markdownv2",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
			{Text: "Click me for help!", Url: fmt.Sprintf("t.me/%s?start=help", b.Username)},
		}}}})
	return err
}
