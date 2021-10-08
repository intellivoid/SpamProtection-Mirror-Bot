package main

import (
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/database"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/handlers"
	"go.uber.org/zap"
	"log"

	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/core"
)

func main() {
	e := core.NewShiinaConfig()
	err := e.ReadFile("spb.conf")
	if err != nil {
		log.Fatal(err)
	}
	loggerMgr := core.InitZapLog(e.Debug)
	zap.ReplaceGlobals(loggerMgr)
	defer func(loggerMgr *zap.Logger) {
		err := loggerMgr.Sync()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(loggerMgr) // flushes buffer, if any
	logger := loggerMgr.Sugar()
	core.SUGARED = loggerMgr.Sugar()
	b, up, err := core.BotInit(e)
	if err != nil {
		log.Fatal(err)
	}
	handlers.LoadHandlers(up.Dispatcher)
	database.StartDatabase(e.DatabaseName)
	logger.Info(b.FirstName, " has started, ID: ", b.Id)
	up.Idle()
}
