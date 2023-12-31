package main

import(
	"log"
	"os"

	"discordGameStarter/bot/commands"
	"github.com/bwmarrin/discordgo"
)


func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	
	defer discord.Close()
	

	log.Println("Adding commands...")
	for _, command := range commands.Commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", command.Name, err)
		}
		log.Printf("Added %s", command.Name)
	}
	log.Println("Commands added successfully")
}
