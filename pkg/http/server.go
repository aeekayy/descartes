package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"github.com/aeekayy/descartes/pkg/config"
)

// NewServer return a new Fiber web server
// for running the server side of Descartes
func NewServer(config *config.AppConfig) (*fiber.App, error) {
	app := fiber.New()

	fmt.Printf("Using port of %s\n", config.Port)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"pong":    true,
		})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"pong":    true,
		})
	})

	return app, nil
}
