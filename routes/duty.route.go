package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func DutyRoute(app *fiber.App) {
	//All routes related to water_areas comes here
	app.Get("/duty/:id", controllers.GetADuty)
	app.Get("/duties", controllers.GetAllDuties)
}
