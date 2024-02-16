package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	crons "github.com/louislaugier/campaign-bot/internal"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		godotenv.Load()
	}

	go crons.Schedule()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	fmt.Println("Server is listening on port 80...")
	http.ListenAndServe(":80", nil)
}
