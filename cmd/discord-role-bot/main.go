package main

import (
	"discord-role-bot/internal/bot"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	file, err := os.OpenFile("logs/discord-role-bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal("error opening log file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	log.SetOutput(file)

	b := bot.Bot
	b.Start()
	defer b.DiscordSession.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop
	log.Print("shutting down")
}
