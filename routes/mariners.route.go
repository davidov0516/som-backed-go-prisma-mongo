package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func MarinersRoute(app *fiber.App) {
	//All routes related to mariners comes here
	app.Post("/mariner", controllers.CreateMariner)
	app.Get("/mariner/:id", controllers.GetAMariner)
	app.Put("/mariner/:id", controllers.EditAMariner)
	app.Delete("/mariner/:id", controllers.DeleteAMariner)
	app.Get("/mariners", controllers.GetAllMariners)
}
