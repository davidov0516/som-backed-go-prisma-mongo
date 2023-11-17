package main

import (
	"som-backend/configs"
	"som-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// cors problem fix
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// connect to db (prepare with prisma)
	configs.ConnectPrismaClient()

	// Defer the disconnect function to be called when the program ends
	defer func() {
		configs.DisconnectPrismaClient()
	}()

	// setup routes
	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & prisma & mongoDB"})
	})

	// run program
	app.Listen(":6060")
}
