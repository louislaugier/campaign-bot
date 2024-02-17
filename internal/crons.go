package crons

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	brevo "github.com/louislaugier/campaign-bot/internal/pkg"
)

var executedToday bool

func Schedule() {
	// location := time.FixedZone("UTC+5", 5*60*60)  // UTC +5 hours
	location := time.FixedZone("UTC+5:30", 5*60*60+30*60) // UTC +5 hours 30 minutes

	for {
		now := time.Now().In(location)
		isEightOClock := now.Hour() == 8 && now.Minute() == 00

		if !isEightOClock && !executedToday {
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
	log.Println("Sending all campaigns now.")

	accountsCount, _ := strconv.Atoi(os.Getenv("ACCOUNTS_COUNT"))

	for i := 1; i <= accountsCount; i++ {
		apiKey := os.Getenv(fmt.Sprintf("KEY%d", i))
		cl := brevo.NewBrevoClient(apiKey)

		err := brevo.SendCampaign(cl)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	log.Println("Campaigns sent successfully.")
}
