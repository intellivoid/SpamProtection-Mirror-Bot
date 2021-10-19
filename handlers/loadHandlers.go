package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadHandlers(d *ext.Dispatcher) {
	startCMD := handlers.NewCommand(StartCmd, startHandler)
	helpCMD := handlers.NewCommand(HelpCmd, helpHandler)
	getMeCMD := handlers.NewCommand(FetchCmd, fetchHandler)
	settingsCMD := handlers.NewCommand(SettingsCmd, settingsHandler)
	autoBan := handlers.NewMessage(msgFilter, autoBanHandler)
	logHandler := handlers.NewMessage(logChatFilter, logChat)
	d.AddHandler(startCMD)
	d.AddHandler(helpCMD)
	d.AddHandler(getMeCMD)
	d.AddHandler(settingsCMD)
	d.AddHandler(autoBan)
	d.AddHandler(logHandler)
}
