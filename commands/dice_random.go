package commands

import (
	"github.com/bufferof/bot/response"
	"github.com/bwmarrin/discordgo"
)

var DiceRandom = &discordgo.ApplicationCommand{
	Name: "랜덤 주사위",
	Description: "주사위를 굴려서 숫자를 맞추면 상금을 줍니다",
}

func DiceRandomHandler(s *discordgo.Session, i *discordgo.InteractionCreate){
	
}