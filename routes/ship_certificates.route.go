package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Ship_CertificateRoute(app *fiber.App) {
	//All routes related to charaters comes here
	app.Post("/ship_certificate", controllers.CreateShipCertificate)
	app.Get("/ship_certificate/:id", controllers.GetAShipCertificate)
	app.Put("/ship_certificate/:id", controllers.EditAShipCertificate)
	app.Delete("/ship_certificate/:id", controllers.DeleteAShipCertificate)
	app.Get("/ship_certificates", controllers.GetAllShipCertificates)
}
