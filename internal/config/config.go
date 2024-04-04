package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	DiscordToken string
	GuildID      string
	RoleConfig   map[string]map[string]string
}

func (c *Config) LoadEnv() {
	configContents, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		configContents, err = ioutil.ReadFile(
			"/etc/discord-role-bot/config.json",
		)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	err = json.Unmarshal(configContents, &c)
	if err != nil {
		log.Fatal("error during json unmarshal: ", err)
	}
}
