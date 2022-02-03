package main

import (
	"the-truth-about/api"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	app := fiber.New()
	api.StartUserFlow(app)
	app.Listen(":3000")
}
