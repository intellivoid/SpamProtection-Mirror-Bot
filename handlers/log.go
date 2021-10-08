package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/core"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/database"
)

func logChat(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	d, err := database.GetChat(chat.Id)
	if err != nil {
		core.SUGARED.Error(err)
		return err
	}
	if d.ChatID == 0 {
		database.InsertChat(chat.Id, false, false, "")
	}
	// fmt.Println(d)
	return ext.ContinueGroups
}

func logChatFilter(msg *gotgbot.Message) bool {
	return msg.Chat.Type != "private"
}
