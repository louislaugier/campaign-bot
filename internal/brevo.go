package brevo

import (
	"context"
	"os"

	brevo "github.com/getbrevo/brevo-go/lib"
)

func NewBrevoClient(apiKey string) *brevo.APIClient {
	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", os.Getenv(apiKey))

	br := brevo.NewAPIClient(cfg)

	return br
}

func SendCampaign(b *brevo.APIClient) error {
	_, err := b.EmailCampaignsApi.SendEmailCampaignNow(context.Background(), int64(1))
	if err != nil {
		return err
	}

	return nil
}
