package handler

import (
	"github.com/bufferof/bot/commands"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	name := i.ApplicationCommandData().Name

	if handler, ok := commands.Handlers[name]; ok {
		handler(s, i)
	}
}
