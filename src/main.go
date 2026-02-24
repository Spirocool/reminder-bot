package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load .env file: %s", err)
	}

	TOKEN := os.Getenv("TOKEN")

	discord, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatalf("Couldn't create discord session: %s", err)
	}

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := handlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	err = discord.Open()
	if err != nil {
		log.Fatalf("Couldn't start discord session: %s", err)
	}

	fmt.Println("Creating commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	fmt.Println("Bot is online...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// registeredCommands, err = discord.ApplicationCommands(discord.State.User.ID, "")
	// if err != nil {
	// 	log.Fatalf("Couldn't fetch registered commands: %v", err)
	// }

	fmt.Printf("Removing Commands...")
	for _, v := range registeredCommands {
		err := discord.ApplicationCommandDelete(discord.State.User.ID, "", v.ID)
		if err != nil {
			log.Fatalf("Could't delete command '%v': %v", v.Name, err)
		}
	}

	err = discord.Close()
	if err != nil {
		log.Fatalf("Couldn't close discord session gracefully: %s", err)
	}
}
