package main

import (
	"fiber-server/database"
	"fiber-server/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	// "context"
	// "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
	// "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	database.ConnectionDB()

	sqlDb, err := database.DBConn.DB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer sqlDb.Close()

	// // Load AWS Config
	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	log.Fatalf("Failed to load AWS config: %v", err)
	// }

	// // Create S3 Client
	// s3Client := s3.NewFromConfig(cfg)
	// uploader := manager.NewUploader(s3Client)

	// Initialize Fiber app
	app := fiber.New()
	app.Use(logger.New())

	// Setup routes
	routers.SetUpRouter(app)

	// Start server
	log.Fatal(app.Listen(":4000"))
}
