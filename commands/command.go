package commands

import "github.com/bwmarrin/discordgo"

var CmdList = []*discordgo.ApplicationCommand{
	Ping,
}

var Handlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	"ping": PingHandler,
}