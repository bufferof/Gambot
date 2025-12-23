package commands

import (
	"github.com/bufferof/bot/response"
	"github.com/bwmarrin/discordgo"
)

var Ping = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Pong!",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response.Reply(
		s,
		i,
		&discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				response.Make_embed("Pong!", "🏓 Poing~~", 0),
			},
		},
	)
}
