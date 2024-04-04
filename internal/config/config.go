package config

import (
	"log"
	"os"
)

type Config struct {
	DiscordToken string
	GuildID      string
}

func (c *Config) LoadEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Fatal(err)
		os.Exit(1)
	}

	c.DiscordToken = os.Getenv("DISCORD_TOKEN")
	c.GuildID = os.Getenv("GUILD_ID")
}
