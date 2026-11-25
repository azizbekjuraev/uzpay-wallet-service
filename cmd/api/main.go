package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	userapp "uzpay/internal/application/user"
	walletapp "uzpay/internal/application/wallet"
	"uzpay/internal/config"
	httpapi "uzpay/internal/http"
	"uzpay/internal/infrastructure/postgres"
	"uzpay/internal/infrastructure/security"
)

func main() {
	if err := run(); err != nil {
		log.Println("error", err)
	}
}

func run() error {
	ctx := context.Background()

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbpool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// users
	repo := postgres.NewUserRepository(dbpool)
	hasher := security.NewBcryptHasher(12)
	svc := userapp.NewService(repo, hasher)
	mux := http.NewServeMux()
	userHandler := httpapi.NewUserHandler(svc)
	mux.HandleFunc("POST /users", userHandler.Register)

	// wallets
	walletrepo := postgres.NewWalletRepository(dbpool)
	walletsvc := walletapp.NewService(walletrepo)
	walletHandler := httpapi.NewWalletHandler(walletsvc)
	mux.HandleFunc("POST /wallets", walletHandler.Register)

	return http.ListenAndServe(":8080", mux)
}
