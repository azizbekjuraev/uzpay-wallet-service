package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"uzpay/internal/app"
	"uzpay/internal/config"
)

func main () {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	dbpool, err := app.InitPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	log.Println("Listening on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err);
	}
};

