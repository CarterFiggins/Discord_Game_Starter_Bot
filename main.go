package main

import(
	"log"
	"os"

	"discordGameStarter/mongo"
	"discordGameStarter/bot"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	mongo.Connect()
	defer mongo.Close()

	bot.BotToken = botToken
	bot.Run()
}