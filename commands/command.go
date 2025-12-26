package commands

import "github.com/bwmarrin/discordgo"

var CmdList = []*discordgo.ApplicationCommand{
	Ping,
	Slot,
}

var Handlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	"ping": PingHandler,
	"슬롯머신": SlotMachineHandler,
}

var BtnHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	"slot_run": SlotRunHandler,
}
