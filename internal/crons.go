package crons

import (
	"log"
	"os"
	"strings"
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

		if isEightOClock && !executedToday {
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
	sendCampaign := func(APIKey string) error {
		cl := brevo.NewBrevoClient(APIKey)

		err := brevo.SendCampaign(cl)
		if err != nil {
			return err
		}

		return nil
	}

	log.Println("Sending all campaigns now.")

	keysEnv := os.Getenv("KEYS")

	keys := strings.Split(keysEnv, ",")

	if len(keys) > 0 {
		for _, key := range keys {
			if err := sendCampaign(key); err != nil {
				log.Println(err)
				continue
			}
		}
	} else if keysEnv != "" {
		if err := sendCampaign(keysEnv); err != nil {
			log.Println(err)
		}
	}

	log.Println("Campaigns sent successfully.")
}
