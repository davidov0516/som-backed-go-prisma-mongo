package routes

import (
	"som-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Mariner_CertificateRoute(app *fiber.App) {
	//All routes related to charaters comes here
	app.Post("/mariner_certificate", controllers.CreateMarinerCertificate)
	app.Get("/mariner_certificate/:id", controllers.GetAMarinerCertificate)
	app.Put("/mariner_certificate/:id", controllers.EditAMarinerCertificate)
	app.Delete("/mariner_certificate/:id", controllers.DeleteAMarinerCertificate)
	app.Get("/mariner_certificates", controllers.GetAllMarinerCertificates)
}
