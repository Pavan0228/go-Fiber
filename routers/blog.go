package routers

import (
	"fiber-server/controller"

	// "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App ) {


	app.Get("/" , controller.BlogList)

	app.Post("/" , controller.CreateBlog)


	app.Post("/upload" , controller.CompressImg)

	// app.Put("/:id" , controller.UpdateBlog)

	// app.Delete("/:id" , controller.DeleteBlog)

}