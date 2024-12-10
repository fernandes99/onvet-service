package main

import (
	"fmt"
	"log"
	UserRouter "onvet/internal/app/user/routes"
	"onvet/internal/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, world!")

	db.StartConnection()
	defer db.Close()

	app := fiber.New()

	UserRouter.RegisterRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
