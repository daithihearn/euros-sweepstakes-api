// @title Euros Sweepstakes API
// @version 0.1.0
// @description Returns scores for the Euro 2024 sweepstakes
// @BasePath /api/v1
package main

import (
	"context"
	_ "euros-sweepstakes-api/docs"
	"euros-sweepstakes-api/pkg/cache"
	"euros-sweepstakes-api/pkg/score"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"log"
	"net/url"
	"os"
	"strings"

	ginSwagger "github.com/swaggo/gin-swagger"
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
	scoreHandler := score.Handler{S: &scoreService}

	// Set up the API routes.
	router := gin.Default()

	// Configure CORS with custom settings
	// Get the environment variable
	origins := os.Getenv("CORS_ALLOWED_ORIGINS")

	// Check if the environment variable is empty and set a default value
	if origins == "" {
		origins = "http://localhost:888,http://localhost:3000" // Replace with your default list
	}

	config := cors.Config{
		AllowOrigins:  strings.Split(origins, ","),
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}
	router.Use(cors.New(config))

	// Redirect from root to /swagger/index.html
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	// Configure the routes
	router.GET("/api/v1/scores", scoreHandler.Get)

	// Use the generated docs in the docs package.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = router.Run(":" + port)
	if err != nil {
		return
	}

	// Wait for the cancellation of the context (due to signal handling)
	<-ctx.Done()
}
