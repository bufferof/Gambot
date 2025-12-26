package commands

import (
	"github.com/bufferof/bot/response"
	"github.com/bwmarrin/discordgo"
)

var Slot = &discordgo.ApplicationCommand{
	Name:        "슬롯머신",
	Description: "같은 문양이 3개 연속해서 나오면 상금을 얻습니다",
}

func SlotMachineHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response.Reply(
		s,
		i,
		&discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				response.Make_embed("슬롯머신", "아래 버튼을 눌러 스롯머신을 돌리세요!", 0xFFD700),
			},
			Components: []discordgo.MessageComponent{
				response.Make_btn("💰", "slot_run", discordgo.PrimaryButton),
			},
		},
	)
}

func SlotRunHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "test222222",
		},
	})
}
