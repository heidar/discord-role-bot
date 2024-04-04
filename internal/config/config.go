package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	DiscordToken string
	GuildID      string
}

func (c *Config) LoadEnv() {
	configContents, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		log.Fatal("error when opening config file: ", err)
	}

	err = json.Unmarshal(configContents, &c)
	if err != nil {
		log.Fatal("error during json unmarshal: ", err)
	}
	log.Print(c.DiscordToken)
	log.Print(c.GuildID)
}
