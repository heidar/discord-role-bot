package bot

import (
	"discord-role-bot/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Bot bot

type bot struct {
	DiscordSession *discordgo.Session
	Guild          *discordgo.Guild
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
	guild, err := b.DiscordSession.Guild(b.Config.GuildID)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	b.Guild = guild
}

func reactionAddHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
}

func addRole(userID string, roleName string) {
	role, err := findRole(roleName)
	if err != nil {
		log.Print(err)
	}

	err = Bot.DiscordSession.GuildMemberRoleAdd(Bot.Config.GuildID, userID, role.ID)
	if err != nil {
		log.Print(err)
	}
}

func findRole(name string) (*discordgo.Role, error) {
	for _, role := range Bot.Guild.Roles {
		if role.Name == name {
			return role, nil
		}
	}

	return nil, fmt.Errorf("could not find role with name: %s", name)
}
