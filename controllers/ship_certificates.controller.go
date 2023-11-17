package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"som-backend/configs"
	"som-backend/prisma/db"
	"som-backend/responses"
	"time"

	"github.com/gofiber/fiber/v2"
)

// functions
func CreateShipCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient
	fmt.Println("______")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var ship_certificate db.ShipCertificatesModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&ship_certificate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a ship_certificate
	putin_date, ok := ship_certificate.PutinDate()
	if !ok {
		putin_date = time.Now() // Set a default value if phone number is not available
	}

	cert_ID, ok := ship_certificate.CertID()
	if !ok {
		cert_ID = "" // Set a default value if phone number is not available
	}

	issue_date, ok := ship_certificate.IssueDate()
	if !ok {
		issue_date = time.Now() // Set a default value if phone number is not available
	}

	expire_date, ok := ship_certificate.ExpireDate()
	if !ok {
		expire_date = time.Now() // Set a default value if phone number is not available
	}

	account, ok := ship_certificate.Account()
	if !ok {
		account = "" // Set a default value if phone number is not available
	}

	price, ok := ship_certificate.Price()
	if !ok {
		price = 1 // Set a default value if phone number is not available
	}

	reg_fee, ok := ship_certificate.RegFee()
	if !ok {
		reg_fee = 1 // Set a default value if phone number is not available
	}
	fmt.Println("_____")
	createdCertificate, err := prisma.ShipCertificates.CreateOne(
		db.ShipCertificates.Certificate.Link(
			db.Certificates.ID.Equals(ship_certificate.CertificateID),
		),
		db.ShipCertificates.Department.Link(
			db.Departments.ID.Equals(ship_certificate.DepartmentID),
		),
		db.ShipCertificates.Ship.Link(
			db.Ships.ID.Equals(ship_certificate.ShipID),
		),
		db.ShipCertificates.PutinDate.Set(putin_date),
		db.ShipCertificates.CertID.Set(cert_ID),
		db.ShipCertificates.IssueDate.Set(issue_date),
		db.ShipCertificates.ExpireDate.Set(expire_date),
		db.ShipCertificates.Account.Set(account),
		db.ShipCertificates.Price.Set(price),
		db.ShipCertificates.RegFee.Set(reg_fee),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdCertificate, "", "  ")
	fmt.Printf("created ship_certificate: %s\n", result)

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": createdCertificate}})
}

func GetAShipCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	ship_certificate, err := prisma.ShipCertificates.FindUnique(db.ShipCertificates.ID.Equals(id)).With(
		db.ShipCertificates.Ship.Fetch(),
		db.ShipCertificates.Department.Fetch(),
		db.ShipCertificates.Certificate.Fetch(),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": ship_certificate}})
}

func EditAShipCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")

	var ship_certificate db.ShipCertificatesModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&ship_certificate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a ship_certificate
	putin_date, ok := ship_certificate.PutinDate()
	if !ok {
		putin_date = time.Now() // Set a default value if phone number is not available
	}

	cert_ID, ok := ship_certificate.CertID()
	if !ok {
		cert_ID = "" // Set a default value if phone number is not available
	}

	issue_date, ok := ship_certificate.IssueDate()
	if !ok {
		issue_date = time.Now() // Set a default value if phone number is not available
	}

	expire_date, ok := ship_certificate.ExpireDate()
	if !ok {
		expire_date = time.Now() // Set a default value if phone number is not available
	}

	account, ok := ship_certificate.Account()
	if !ok {
		account = "" // Set a default value if phone number is not available
	}

	price, ok := ship_certificate.Price()
	if !ok {
		price = 1 // Set a default value if phone number is not available
	}

	reg_fee, ok := ship_certificate.RegFee()
	if !ok {
		reg_fee = 1 // Set a default value if phone number is not available
	}
	updateCertificate, err := prisma.ShipCertificates.FindUnique(
		db.ShipCertificates.ID.Equals(id),
	).Update(
		db.ShipCertificates.Certificate.Link(
			db.Certificates.ID.Equals(ship_certificate.CertificateID),
		),
		db.ShipCertificates.Department.Link(
			db.Departments.ID.Equals(ship_certificate.DepartmentID),
		),
		db.ShipCertificates.Ship.Link(
			db.Ships.ID.Equals(ship_certificate.ShipID),
		),
		db.ShipCertificates.PutinDate.Set(putin_date),
		db.ShipCertificates.CertID.Set(cert_ID),
		db.ShipCertificates.IssueDate.Set(issue_date),
		db.ShipCertificates.ExpireDate.Set(expire_date),
		db.ShipCertificates.Account.Set(account),
		db.ShipCertificates.Price.Set(price),
		db.ShipCertificates.RegFee.Set(reg_fee),
	).Exec(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateCertificate}})
}

func DeleteAShipCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	_, err := prisma.ShipCertificates.FindUnique(db.ShipCertificates.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Certificate successfully deleted!"}})
}

func GetAllShipCertificates(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ship_certificates, err := prisma.ShipCertificates.FindMany().With(
		db.ShipCertificates.Ship.Fetch(),
		db.ShipCertificates.Department.Fetch(),
		db.ShipCertificates.Certificate.Fetch(),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": ship_certificates}})
}
