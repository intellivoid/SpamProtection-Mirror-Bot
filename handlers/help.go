package handlers

import (
	"fmt"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/helpers"
)

func getHelp(u *gotgbot.User) string {
	msg := mdparser.GetBold("Hai ").AppendMention(u.FirstName, u.Id).AppendBold(", welcome!").AppendNormal("\n\n")
	msg = msg.AppendNormal("I am a bot, with the primary purpose of adding more overhead on @SpamProtectionBot.").AppendNormal("\n")
	msg = msg.AppendNormal("Due to this, I function quite similarly to it, I might look like a clone of it.").AppendNormalThis("\n\n")
	msg = msg.AppendBold("Here are the commands I currently accept").AppendNormal(":\n")
	msg = msg.AppendNormal("- ").AppendMono(StartCmd).AppendNormal(": starts me, as usual\n")
	msg = msg.AppendNormal("- ").AppendMono(HelpCmd).AppendNormal(": makes me send THIS message\n")
	msg = msg.AppendNormal("- ").AppendMono(FetchCmd).AppendNormalThis(": fetch a user on the API")
	return msg.ToString()
}

func helpHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	m := ctx.EffectiveMessage
	if m.Chat.Type != "private" {
		_, err := m.Reply(b, "Contact me in PM", &gotgbot.SendMessageOpts{ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
			{Text: "Click me for help!", Url: fmt.Sprintf("t.me/%s?start=help", b.Username)},
		}}}})
		if err != nil {
			return helpers.SendError(err, ctx, b)
		}

	} else {
		_, err := b.SendMessage(m.Chat.Id, getHelp(m.From), &gotgbot.SendMessageOpts{ParseMode: "markdownv2"})
		return err
	}
	return nil
}
