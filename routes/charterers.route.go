package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func CharterersRoute(app *fiber.App) {
	//All routes related to charaters comes here
	app.Post("/charterer", controllers.CreateCharterer)
	app.Get("/charterer/:id", controllers.GetACharterer)
	app.Put("/charterer/:id", controllers.EditACharterer)
	app.Delete("/charterer/:id", controllers.DeleteACharterer)
	app.Get("/charterers", controllers.GetAllCharterers)
}
