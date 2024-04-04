package bot

import (
	"discord-role-bot/internal/config"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Bot bot

type bot struct {
	DiscordSession *discordgo.Session
	Config         config.Config
}

func (b *bot) initialize() {
	b.Config.LoadEnv()
}

func (b *bot) Start() {
	b.initialize()

	discordSession, err := discordgo.New("Bot " + b.Config.DiscordToken)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = discordSession.Open()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	b.DiscordSession = discordSession
}
