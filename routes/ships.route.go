package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ShipsRoute(app *fiber.App) {
	//All routes related to ships comes here
	app.Post("/ship", controllers.CreateShip)
	app.Get("/ship/:id", controllers.GetAShip)
	app.Put("/ship/:id", controllers.EditAShip)
	app.Delete("/ship/:id", controllers.DeleteAShip)
	app.Get("/ships", controllers.GetAllShips)
}
