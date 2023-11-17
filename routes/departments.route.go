package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func DepartmentsRoute(app *fiber.App) {
	//All routes related to charaters comes here
	app.Post("/department", controllers.CreateDepartment)
	app.Get("/department/:id", controllers.GetADepartment)
	app.Put("/department/:id", controllers.EditADepartment)
	app.Delete("/department/:id", controllers.DeleteADepartment)
	app.Get("/departments", controllers.GetAllDepartments)
}
