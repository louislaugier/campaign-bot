package main

import (
	"os"

	"github.com/joho/godotenv"
	crons "github.com/louislaugier/campaign-bot/internal"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		godotenv.Load()
	}

	crons.Schedule()
}
