package products

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	service := newService()

	rootGroup := app.Group("/api/v1/products")

	rootGroup.Get("/", service.GetProducts)
}