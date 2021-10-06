package core

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"go.uber.org/zap"
	"net/http"
)

func BotInit(config *ShiinaConfig) (b *gotgbot.Bot, updater ext.Updater,  err error) {
	b, err = gotgbot.NewBot(config.BotToken, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		SUGARED.Error(fmt.Sprintf("Failed to create new bot due to %s", err.Error()), zap.Error(err))
		return nil, updater, err
	}
	updater = ext.NewUpdater(nil)
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: config.DropUpdates})
	if err != nil {
		SUGARED.Error(fmt.Sprintf("Failed to start polling due to %s", err.Error()), zap.Error(err))
	}
	SUGARED.Info(fmt.Sprintf("%s has started | ID: %d", b.Username, b.Id))
	return b, updater, nil
}
