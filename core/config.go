package core

import (
	"fmt"
	hocon "github.com/go-akka/configuration"
	"io/ioutil"
)

type ShiinaConfig struct {
	BotToken       string
	DatabaseName   string
	CoffeeHouseKey string
	BotAdmins      []int64
	Debug          bool
	DropUpdates    bool
}

var Data *ShiinaConfig

func NewShiinaConfig() *ShiinaConfig {
	return &ShiinaConfig{}
}

func (c *ShiinaConfig) ReadFile(path string) error {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	conf := hocon.ParseString(string(b))
	// print(conf.String())
	c.BotToken = conf.GetString("ShiinaConfig.BotToken")
	c.DatabaseName = conf.GetString("ShiinaConfig.DatabaseName")
	c.BotAdmins = conf.GetInt64List("ShiinaConfig.BotAdmins")
	c.Debug = conf.GetBoolean("ShiinaConfig.Debug")
	c.DropUpdates = conf.GetBoolean("ShiinaConfig.DropUpdates")
	Data = c
	return nil
}

func (c *ShiinaConfig) GetBotAdmins() []int64 {
	var d []int64
	for i := range c.BotAdmins {
		d = append(d, c.BotAdmins[i])
	}
	return d
}
