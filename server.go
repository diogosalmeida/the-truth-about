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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Person{Name: "Diogo", Email: "diogomean@gmail.com"})
	})

	app.Listen(":3000")
}
