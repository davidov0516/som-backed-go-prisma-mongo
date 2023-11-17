package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	CharterersRoute(app)
	MarinersRoute(app)
	WaterAreasRoute(app)
	CertificateRoute(app)
	Ship_CertificateRoute(app)
	Mariner_CertificateRoute(app)
	ShipsRoute(app)
	ShippingsRoute(app)
	DepartmentsRoute(app)
	DutyRoute(app)
}
