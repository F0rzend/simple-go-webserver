package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/repositories"
	"github.com/F0rzend/simple-go-webserver/app/aggregate/user/repositories"
	"github.com/F0rzend/simple-go-webserver/app/server"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))

	address := getEnv("ADDRESS", ":8080")
	logger.Info("run api", slog.String("address", address))

	userRepository := userrepositories.NewMemoryUserRepository()
	bitcoinRepository := bitcoinrepositories.NewMemoryBTCRepository()

	apiServer := server.NewServer(userRepository, bitcoinRepository, bitcoinRepository)

	srv := &http.Server{
		Addr:              address,
		Handler:           apiServer.GetHTTPHandler(logger),
		ReadHeaderTimeout: 1 * time.Second,
	}
	logger.Error("server stopped", srv.ListenAndServe())
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
