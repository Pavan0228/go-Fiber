package controller

import (
	"fiber-server/database"
	"fiber-server/model"
	"fmt"
	"log"

	// "context"
	// "os"

	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/disintegration/imaging"
)

func BlogList(c *fiber.Ctx) error {

	// Prepare the context with a default success message
	context := fiber.Map{
		"message": "list of blogs",
		"status":  "success",
	}

	// Get the database connection
	db := database.DBConn
	if db == nil {
		// Check if the database connection is not available
		log.Println("Database connection is nil")
		context["message"] = "Database connection failed"
		context["status"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// Define a slice to hold the records
	var records []model.Blog

	// Perform the database query to find all blogs
	if err := db.Find(&records).Error; err != nil {
		// Handle database query error
		log.Printf("Error fetching blog records: %v", err)
		context["message"] = "Failed to retrieve blogs"
		context["status"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// Check if no records were found
	if len(records) == 0 {
		// If no blogs are found, return a message indicating so
		context["message"] = "No blogs found"
		context["status"] = "success"
		context["data"] = []model.Blog{}
		return c.JSON(context)
	}

	// If the query was successful, include the data in the response
	context["data"] = records
	return c.JSON(context)
}


func CreateBlog(c *fiber.Ctx) error {
	// Prepare the default response context
	context := fiber.Map{
		"message": "create blog",
		"status":  "success",
	}

	// Create a new blog record instance
	records := new(model.Blog)

	// Parse the incoming JSON body to the records struct
	if err := c.BodyParser(&records); err != nil {
		// Log the error and send a response
		log.Printf("Error parsing body: %v", err)
		context["message"] = "Error parsing data"
		context["status"] = "error"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Validate the blog data (optional)
	if records.Title == "" || records.Post == "" {
		// Log the missing required fields
		log.Println("Missing required fields: Title or Post")
		context["message"] = "Title and Post are required"
		context["status"] = "error"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	// Get the database connection
	db := database.DBConn
	if db == nil {
		// Check if the database connection is not available
		log.Println("Database connection is nil")
		context["message"] = "Database connection failed"
		context["status"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// Try to create the blog record in the database
	result := db.Create(&records)
	if result.Error != nil {
		// Log the database creation error
		log.Printf("Error creating record: %v", result.Error)
		context["message"] = "Error creating record"
		context["status"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// If creation is successful, return the blog data
	context["data"] = records
	return c.Status(fiber.StatusCreated).JSON(context)
}


// func UpdateBlog(c *fiber.Ctx) error{

// }

// func DeleteBlog(c *fiber.Ctx) error{

// }


func CompressImg(c *fiber.Ctx) error {
	response := fiber.Map{
		"message": "compress image",
		"status":  "success",
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error parsing file: %v", err)
		response["message"] = "Error parsing file"
		response["status"] = "error"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	filePath := "public/uploads/" + file.Filename
	err = c.SaveFile(file, filePath)
	if err != nil {
		log.Printf("Error saving file: %v", err)
		response["message"] = "Error saving file"
		response["status"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	fileContent, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		response["message"] = "Error opening file"
		response["status"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	defer fileContent.Close()

	img, err := imaging.Decode(fileContent)
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		response["message"] = "Error decoding image"
		response["status"] = "error"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	compressedImg := imaging.Resize(img, 800, 0, imaging.Lanczos)

	compressedFilePath := fmt.Sprintf("public/compressed/compressed_%s", file.Filename)

	err = imaging.Save(compressedImg, compressedFilePath, imaging.JPEGQuality(80))
	if err != nil {
		log.Printf("Error saving compressed image: %v", err)
		response["message"] = "Error saving compressed image"
		response["status"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response["message"] = "File uploaded and compressed successfully"
	response["compressedFilePath"] = compressedFilePath
	return c.JSON(response)
}


	// // Upload to S3
	// uploadResult, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String("autofinancetrack-pennytracker"), // Replace with your bucket name
	// 	Key:    aws.String(file.Filename),
	// 	Body:   openedFile,
	// })
	// if err != nil {
	// 	log.Printf("Error uploading to S3: %v", err)
	// 	response["message"] = "Error uploading file"
	// 	response["status"] = "error"
	// 	return c.Status(fiber.StatusInternalServerError).JSON(response)
	// }
