package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ShippingsRoute(app *fiber.App) {
	//All routes related to charaters comes here
	app.Post("/shipping", controllers.CreateShipping)
	app.Get("/shipping/:id", controllers.GetAShipping)
	app.Put("/shipping/:id", controllers.EditAShipping)
	app.Delete("/shipping/:id", controllers.DeleteAShipping)
	app.Get("/shippings", controllers.GetAllShippings)
}
