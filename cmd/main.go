package main

import (
	"github.com/joho/godotenv"
	"github.com/louislaugier/campaign-bot/crons"
)

func main() {
	godotenv.Load()

	crons.Schedule()
}
