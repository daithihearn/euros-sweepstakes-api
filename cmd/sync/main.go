package main

import (
	"context"
	"euros-sweepstakes-api/pkg/cache"
	"euros-sweepstakes-api/pkg/score"
	"euros-sweepstakes-api/pkg/sheets"
	"euros-sweepstakes-api/pkg/sync"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
)

func init() {
	// Load .env file if it exists
	_ = godotenv.Load()
}

func main() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get the Redis URL from the environment
	redisUri := os.Getenv("REDIS_URL")
	if redisUri == "" {
		redisUri = "redis://:password@localhost:6379/0"
	}

	// Parse the Redis URL
	parsedUrl, err := url.Parse(redisUri)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	// Extract the password from the URL
	redisPassword, _ := parsedUrl.User.Password()

	// Extract the address from the URL
	redisAddr := parsedUrl.Host

	// Configure the Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	scoreCache := cache.NewRedisCache[[]score.Score](rdb, ctx)
	scoreService := score.Service{Cache: scoreCache}

	sheetService, err := sheets.NewSheetService(ctx)
	if err != nil {
		log.Fatalf("Failed to create sheet service: %v", err)
	}

	// TODO: Need an Odds service

	syncService := sync.Syncer{ScoreService: &scoreService, SheetService: sheetService}

	// Sync with the API.
	err = syncService.Sync()
	if err != nil {
		cancel()
		log.Fatal("Failed to sync : ", err)
	}

	cancel()

	// Wait for the cancellation of the context (due to signal handling)
	<-ctx.Done()
}
