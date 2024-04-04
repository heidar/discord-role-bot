package main

import (
	"discord-role-bot/internal/bot"
	"log"
	"os"
	"os/signal"
)

func main() {
	b := &bot.Bot
	b.Start()
	defer b.DiscordSession.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop
	log.Print("shutting down")
}
