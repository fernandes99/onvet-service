package router

import (
	"onvet/internal/app/user/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,
) {
	println("Register Router")

	apiUser := app.Group("v1/users")

	apiUser.Get("/", handlers.GetAll)
	apiUser.Post("/", handlers.Create)

	apiUser.Get("/:id", handlers.GetById)

	apiUser.Get("/:id/addresses", handlers.GetUserAddresses)
	apiUser.Post("/:id/addresses", handlers.CreateUserAddress)

	apiUser.Post("/:id/pets", handlers.CreateUserPet)
	// apiUser.Get("/:id/pets", handlers.CreateUserPet)
}
