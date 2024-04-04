package main

import (
	"discord-role-bot/internal/bot"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	file, err := os.OpenFile("discord-role-bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
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
