package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadHandlers(d *ext.Dispatcher) {
	startCMD := handlers.NewCommand(START_CMD, startHandler)
	helpCMD := handlers.NewCommand(HELP_CMD, helpHandler)
	getMeCMD := handlers.NewCommand(FETCH_CMD, fetchHandler)
	autoBan := handlers.NewMessage(msgFilter, autoBanHandler)
	logHandler := handlers.NewMessage(logChatFilter, logChat)
	d.AddHandler(startCMD)
	d.AddHandler(helpCMD)
	d.AddHandler(getMeCMD)
	d.AddHandler(autoBan)
	d.AddHandler(logHandler)
}
