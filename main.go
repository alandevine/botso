package main

import (
	"log"
	"os"

	"adevine/botso/bot"
)

func main() {
	discordApiToken, ok := os.LookupEnv("DISCORD_API_TOKEN")

	if !ok {
		log.Fatalln("Failed to load environment varaible `DISCORD_API_TOKEN`")
	}

	bot.BotToken = discordApiToken
	err := bot.Run()
	if err != nil {
		log.Panicln("Encountered Error during runtime: " + err.Error())
	}
}
