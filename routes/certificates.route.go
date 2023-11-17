package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func CertificateRoute(app *fiber.App) {
	//All routes related to charaters comes here
	app.Post("/certificate", controllers.CreateCertificate)
	app.Get("/certificate/:id", controllers.GetACertificate)
	app.Put("/certificate/:id", controllers.EditACertificate)
	app.Delete("/certificate/:id", controllers.DeleteACertificate)
	app.Get("/certificates", controllers.GetAllCertificates)
}
