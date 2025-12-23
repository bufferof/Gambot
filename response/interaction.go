package response

import "github.com/bwmarrin/discordgo"

func Reply(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	data *discordgo.InteractionResponseData,
) error {
	return  s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	})
}