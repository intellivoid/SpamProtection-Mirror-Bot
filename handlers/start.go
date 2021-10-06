package handlers

import (
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func startHandler(b *gotgbot.Bot, ctx *ext.Context) error {
    txt := mdparser.GetBold("Hi, I'm " + b.FirstName).AppendNormal("\n\n")
	txt = txt.AppendNormal(`I'm a mirror of @SpamProtectionBot, this means`)
	txt = txt.AppendNormal(` I enforce all bans made through that bot and also provide user information related to it`)
	_, err := b.SendMessage(ctx.Message.Chat.Id, txt.ToString(), &gotgbot.SendMessageOpts{ParseMode: "markdownv2"})
	return err
}
