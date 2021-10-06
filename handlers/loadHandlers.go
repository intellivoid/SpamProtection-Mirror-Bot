package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadHandlers(d *ext.Dispatcher) {
	startCMD := handlers.NewCommand("start", startHandler)
	getMeCMD := handlers.NewCommand("fetch", fetchHandler)
	autoBan  := handlers.NewMessage(msgFilter, autoBanHandler)
	d.AddHandler(startCMD)
	d.AddHandler(getMeCMD)
	d.AddHandler(autoBan)
}
