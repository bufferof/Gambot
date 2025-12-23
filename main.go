package main

import (
	"log"

	"github.com/bufferof/bot/commands"
	"github.com/bufferof/bot/config"
	"github.com/bufferof/bot/internal/handler"
	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("[...] Bot now Starting")
	log.Println("[...] Loading Config")

	config.Load()
	log.Println("[+] Config Loaded!")

	token := config.Token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	dg.Identify.Intents = discordgo.IntentsAll

	dg.AddHandler(handler.InteractionCreate)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	for _, cmd := range commands.CmdList {
		_, err = dg.ApplicationCommandCreate(
			dg.State.User.ID,
			config.GuildID,
			cmd,
		)

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("[+] Setup Success~~")

	select {}
}
