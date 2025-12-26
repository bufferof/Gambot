package response

import "github.com/bwmarrin/discordgo"

func Make_embed(title, description string, color int) *discordgo.MessageEmbed {
	if color == 0 {
		color = 0x00D4FF
	}

	return &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       color,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Made with by Codevate",
		},
	}
}

func Error_embed(description string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error",
		Description: description,
		Color:       0xED4245,
	}
}

func Make_btn(label, cID string, style discordgo.ButtonStyle) discordgo.Button{
	return  discordgo.Button{
		Label: label,
		Style: style,
		CustomID: cID,
	}
}
