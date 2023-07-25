package server

import (
	"github.com/gofiber/fiber/v2"
	"net"
)

func EmbedServer(port string) {
	go func() {
		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("ok")
		})

		err := app.Listen(net.JoinHostPort("127.0.0.1", port))
		if err != nil {
			panic("embed server error")
		}
	}()
}
