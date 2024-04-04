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
	b.DiscordSession.AddHandler(reactionRemoveHandler)

	log.Print("ready for reactions")
}

func reactionAddHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	roleID, err := lookupRoleID(m.MessageReaction, "added")
	if err != nil {
		log.Print(err)
	}
	addRole(m.MessageReaction.UserID, roleID)
}

func addRole(userID, roleID string) {
	err := Bot.DiscordSession.GuildMemberRoleAdd(Bot.Config.GuildID, userID, roleID)
	if err != nil {
		log.Print("error adding role to user: ", err)
	}
}

func reactionRemoveHandler(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	roleID, err := lookupRoleID(m.MessageReaction, "removed")
	if err != nil {
		log.Print(err)
	}
	removeRole(m.MessageReaction.UserID, roleID)
}

func removeRole(userID, roleID string) {
	err := Bot.DiscordSession.GuildMemberRoleRemove(Bot.Config.GuildID, userID, roleID)
	if err != nil {
		log.Print("error removing role from user: ", err)
	}
}

func lookupRoleID(reaction *discordgo.MessageReaction, action string) (string, error) {
	userID := reaction.UserID
	messageID := reaction.MessageID
	emoji := reaction.Emoji.Name

	log.Printf("%v reaction: user: %v, message id: %v, emoji: %v",
		action,
		userID,
		messageID,
		emoji)

	if roleID, ok := Bot.Config.RoleConfig[messageID][emoji]; ok {
		return roleID, nil
	} else {
		return "", fmt.Errorf("could not find role id for message id %v and emoji %v", messageID, emoji)
	}
}
