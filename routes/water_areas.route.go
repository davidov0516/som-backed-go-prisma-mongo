package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func WaterAreasRoute(app *fiber.App) {
	//All routes related to water_areas comes here
	app.Post("/water_area", controllers.CreateWaterArea)
	app.Get("/water_area/:id", controllers.GetAWaterArea)
	app.Put("/water_area/:id", controllers.EditAWaterArea)
	app.Delete("/water_area/:id", controllers.DeleteAWaterArea)
	app.Get("/water_areas", controllers.GetAllWaterAreas)
}
