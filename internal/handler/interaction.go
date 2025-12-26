package handler

import (
	"github.com/bufferof/bot/commands"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		name := i.ApplicationCommandData().Name

		if handler, ok := commands.Handlers[name]; ok {
			handler(s, i)
		}

	case discordgo.InteractionMessageComponent:
		cid := i.MessageComponentData().CustomID

		if handler, ok := commands.BtnHandlers[cid]; ok {
			handler(s, i)
		}
	}
}
