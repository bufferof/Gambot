package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token   string
	GuildID string
	DB_Dsn  string
)


func Load() {
	_ = godotenv.Load()

	Token = os.Getenv("BOT_TOKEN")
	GuildID = os.Getenv("GUILD_ID")
	DB_Dsn = os.Getenv("DB_DSN")

	if Token == "" {
		log.Fatal("TOKEN not SET!")
	}
}
