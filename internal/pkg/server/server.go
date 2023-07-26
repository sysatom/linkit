package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sysatom/linkit/internal/pkg/logs"
	"net"
)

func EmbedServer(port string) {
	go func() {
		app := fiber.New()
		app.Use(cors.New())
		app.Use(recover.New())
		app.Use(requestid.New())

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("ok")
		})

		err := app.Listen(net.JoinHostPort("127.0.0.1", port))
		if err != nil {
			logs.Fatal("embed server error")
		}
	}()
}
