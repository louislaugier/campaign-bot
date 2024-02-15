package crons

import (
	"fmt"
	"os"
	"strconv"
	"time"

	brevo "github.com/louislaugier/campaign-bot/internal"
)

var executedToday bool

func Schedule() {
	// location := time.FixedZone("UTC+5", 5*60*60)  // UTC +5 hours
	location := time.FixedZone("UTC+5:30", 5*60*60+30*60) // UTC +5 hours 30 minutes

	go sendCampaigns()

	for {
		now := time.Now().In(location)
		isExactlyEightOClock := now.Hour() == 8 && now.Minute() == 0 && now.Second() == 0

		if isExactlyEightOClock && !executedToday {
			executedToday = true
			go sendCampaigns()
		}

		if now.Hour() != 8 || now.Minute() != 0 {
			executedToday = false
		}

		time.Sleep(58 * time.Second)
	}
}

func sendCampaigns() {
	accountsCount, _ := strconv.Atoi(os.Getenv("ACCOUNTS_COUNT"))

	for i := 0; i < accountsCount; i++ {
		apiKey := os.Getenv(fmt.Sprintf("KEY%d", i))

		cl := brevo.NewBrevoClient(apiKey)

		brevo.SendCampaign(cl)
	}
}
