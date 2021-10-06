package handlers

import (
	"errors"
	"fmt"
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/Intellivoid/Intellivoid.SpamProtection-go/spamProtection"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/core"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/helpers"
	"strconv"
)

func fetchHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	var data *spamProtection.APIResponse
	UserID, err := helpers.GetID(ctx)
	if err != nil {
		data, err = spamProtection.GetInfoByUsername(ctx.Args()[1])
		if err != nil {
			helpers.SendError(err, ctx, b)
			return err
		}
	} else {
		data, err = spamProtection.GetInfoByID(UserID)
	}
	if err != nil {
		core.SUGARED.Error(err)
	}
	if !data.Success {
		helpers.SendError(errors.New(data.Error.Message), ctx, b)
		return nil
	}
	//var msg mdparser.WMarkDown
	var ct string
	if data.IsUser() {
		ct = "User"
	} else {
		// msg = msg.AppendBold("Chat Information").AppendNormal("\n\n")
		ct = "Chat"
	}
	msg := mdparser.GetBold(fmt.Sprintf("%s Information", ct)).AppendNormal("\n\n")
	if data.IsVerified() {
		msg = msg.AppendBold(fmt.Sprintf("✅ This %s's Telegram account is verified by Intellivoid Accounts", ct)).AppendNormal("\n")
	}
	if data.IsOperator() {
		msg = msg.AppendBold(fmt.Sprintf("👮 This %s is an operator who can blacklist users", ct)).AppendNormal("\n")
	}
	if data.IsAgent() {
		msg = msg.AppendBold(fmt.Sprintf("👮 This %s is an agent who actively reports spam automatically", ct)).AppendNormal("\n")
	}
	if data.IsBlacklisted() {
		msg = msg.AppendBold(fmt.Sprintf("⚠ This %s is blacklisted!", ct)).AppendNormal("\n")
	}
	if data.IsVerified() || data.IsOperator() || data.IsAgent() || data.IsAgent() || data.IsBlacklisted() {
		msg = msg.AppendNormal("\n")
	}
	user, _ := b.GetChat(UserID)
	msg = msg.AppendBold("Private ID").AppendNormal(":\n").AppendMono(data.Results.PrivateTelegramID).AppendNormal("\n")
	if UserID != 0 {
		msg = msg.AppendBold("User ID").AppendNormal(": ").AppendMono(strconv.FormatInt(UserID, 10)).AppendNormal("\n")
		msg = msg.AppendBold("First Name").AppendNormal(": ").AppendMono(user.FirstName).AppendNormal("\n")
		if user.LastName != "" {
			msg = msg.AppendBold("Last Name").AppendNormal(": ").AppendMono(user.LastName).AppendNormal("\n")
		}
		if user.Username != "" {
			msg = msg.AppendBold("Username").AppendNormal(": ").AppendMono(user.Username).AppendNormal("\n")
		}
	}
	if r := data.Results.SpamPrediction; r != nil {
		msg = msg.AppendBold("Spam Prediction").AppendNormal(": ").AppendMono(fmt.Sprintf("%f", r.SpamPrediction)).AppendNormal("\n")
		msg = msg.AppendBold("Ham Prediction").AppendNormal(": ").AppendMono(fmt.Sprintf("%f", r.HamPrediction)).AppendNormal("\n")
	}
	if l := data.Results.LanguagePrediction; l != nil {
		msg = msg.AppendBold("Language").AppendNormal(": ").AppendMono(fmt.Sprintf("%s (%f)", l.Language, l.Probability)).AppendNormal("\n")
	}
	if UserID != 0 {
		msg = msg.AppendBold("Permalink").AppendNormal(": ").AppendMention("here", UserID)
	}
	_, err = b.SendMessage(chat.Id, msg.ToString(), &gotgbot.SendMessageOpts{ParseMode: "markdownv2"})
	if err != nil {
		return err
	}
	return nil
}
