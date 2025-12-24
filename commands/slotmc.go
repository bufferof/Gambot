package commands

import (
	"github.com/bufferof/bot/response"
	"github.com/bwmarrin/discordgo"
)

var Slot = &discordgo.ApplicationCommand{
	Name:        "슬롯머신",
	Description: "같은 문양이 3개 연속해서 나오면 상금을 얻습니다",
}

func SlotMachineHandler(s *discordgo.Session, i *discordgo.InteractionCreate){
	
}