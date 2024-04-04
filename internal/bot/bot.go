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
		log.Fatal("error creating discord sessions: ", err)
		os.Exit(1)
	}

	err = discordSession.Open()
	if err != nil {
		log.Fatal("error opening discord session: ", err)
		os.Exit(1)
	}

	b.DiscordSession = discordSession
	guild, err := b.DiscordSession.Guild(b.Config.GuildID)
	if err != nil {
		log.Fatal("error looking up discord guild: ", err)
		os.Exit(1)
	}
	b.Guild = guild
	b.DiscordSession.AddHandler(reactionAddHandler)

	log.Print("ready for reactions")
}

func reactionAddHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	userID := m.MessageReaction.UserID
	messageID := m.MessageReaction.MessageID
	emoji := m.MessageReaction.Emoji.Name

	log.Printf("added reaction: user: %v, message id: %v, emoji: %v",
		userID,
		messageID,
		emoji)
	roleID := Bot.Config.RoleConfig[messageID][emoji]
	addRole(userID, roleID)
}

func addRole(userID, roleID string) {
	err := Bot.DiscordSession.GuildMemberRoleAdd(Bot.Config.GuildID, userID, roleID)
	if err != nil {
		log.Print("error adding role to user: ", err)
	}
}
