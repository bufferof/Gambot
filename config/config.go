package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


var (
	Token string
	GuildID string
)

func Load() {
	_ = godotenv.Load()

	Token = os.Getenv("BOT_TOKEN")
	GuildID = os.Getenv("GUILD_ID")

	if Token == ""{
		log.Fatal("TOKEN not SET!")
	}
}