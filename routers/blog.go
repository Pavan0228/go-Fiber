package routers

import (
	"fiber-server/controller"

	// "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(app *fiber.App , uploader *manager.Uploader) {


	app.Get("/" , controller.BlogList)

	app.Post("/" , controller.CreateBlog)


	app.Post("/upload", func(c *fiber.Ctx) error {
		return controller.CompressImg(c, uploader)
	})

}