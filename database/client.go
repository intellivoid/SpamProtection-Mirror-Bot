package database

import (
	"fmt"
	"gitlab.com/Dank-del/SpamProtection-Mirror-Bot/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SESSION *gorm.DB

func StartDatabase(name string) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", name)), &gorm.Config{})
	if err != nil {
		core.SUGARED.Error("failed to connect database")
	}

	SESSION = db
	core.SUGARED.Info("Database connected")

	// Create tables if they don't exist
	err = SESSION.AutoMigrate(&Chat{})
	if err != nil {
		core.SUGARED.Error(err)
	}

	core.SUGARED.Info("Auto-migrated database schema")

}
