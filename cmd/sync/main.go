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

	// Configure redis
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = "localhost:6379"
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		redisPassword = "password"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl, // use your Redis Address
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
