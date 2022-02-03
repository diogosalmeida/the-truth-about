package api

import (
	"the-truth-about/models/requests"

	"github.com/gofiber/fiber/v2"
)

// StartUserFlow start flow get user
func StartUserFlow(app *fiber.App) {
	api := app.Group("/users")
	v1 := api.Group("/search")

	v1.Get("/", func(c *fiber.Ctx) error {
		user := requests.UserSearch{
			Name: c.Query("user"),
		}
		return c.JSON(user)
	})

}
