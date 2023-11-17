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
func CreateMarinerCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var mariner_certificate db.MarinerCertificatesModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&mariner_certificate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a mariner_certificate
	putin_date, ok := mariner_certificate.PutinDate()
	if !ok {
		putin_date = time.Now() // Set a default value if phone number is not available
	}

	cert_ID, ok := mariner_certificate.CertID()
	if !ok {
		cert_ID = "" // Set a default value if phone number is not available
	}

	issue_date, ok := mariner_certificate.IssueDate()
	if !ok {
		issue_date = time.Now() // Set a default value if phone number is not available
	}

	expire_date, ok := mariner_certificate.ExpireDate()
	if !ok {
		expire_date = time.Now() // Set a default value if phone number is not available
	}

	account, ok := mariner_certificate.Account()
	if !ok {
		account = "" // Set a default value if phone number is not available
	}

	price, ok := mariner_certificate.Price()
	if !ok {
		price = 1 // Set a default value if phone number is not available
	}

	reg_fee, ok := mariner_certificate.RegFee()
	if !ok {
		reg_fee = 1 // Set a default value if phone number is not available
	}
	createdCertificate, err := prisma.MarinerCertificates.CreateOne(
		db.MarinerCertificates.Certificate.Link(
			db.Certificates.ID.Equals(mariner_certificate.CertificateID),
		),
		db.MarinerCertificates.Department.Link(
			db.Departments.ID.Equals(mariner_certificate.DepartmentID),
		),
		db.MarinerCertificates.Mariner.Link(
			db.Mariners.ID.Equals(mariner_certificate.MarinerID),
		),
		db.MarinerCertificates.Ship.Link(
			db.Ships.ID.Equals(mariner_certificate.ShipID),
		),
		db.MarinerCertificates.PutinDate.Set(putin_date),
		db.MarinerCertificates.CertID.Set(cert_ID),
		db.MarinerCertificates.IssueDate.Set(issue_date),
		db.MarinerCertificates.ExpireDate.Set(expire_date),
		db.MarinerCertificates.Account.Set(account),
		db.MarinerCertificates.Price.Set(price),
		db.MarinerCertificates.RegFee.Set(reg_fee),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdCertificate, "", "  ")
	fmt.Printf("created mariner_certificate: %s\n", result)

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": createdCertificate}})
}

func GetAMarinerCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	mariner_certificate, err := prisma.MarinerCertificates.FindUnique(db.MarinerCertificates.ID.Equals(id)).With(
		db.MarinerCertificates.Ship.Fetch(),
		db.MarinerCertificates.Department.Fetch(),
		db.MarinerCertificates.Certificate.Fetch(),
		db.MarinerCertificates.Mariner.Fetch(),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": mariner_certificate}})
}

func EditAMarinerCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")

	var mariner_certificate db.MarinerCertificatesModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&mariner_certificate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a mariner_certificate
	putin_date, ok := mariner_certificate.PutinDate()
	if !ok {
		putin_date = time.Now() // Set a default value if phone number is not available
	}

	cert_ID, ok := mariner_certificate.CertID()
	if !ok {
		cert_ID = "" // Set a default value if phone number is not available
	}

	issue_date, ok := mariner_certificate.IssueDate()
	if !ok {
		issue_date = time.Now() // Set a default value if phone number is not available
	}

	expire_date, ok := mariner_certificate.ExpireDate()
	if !ok {
		expire_date = time.Now() // Set a default value if phone number is not available
	}

	account, ok := mariner_certificate.Account()
	if !ok {
		account = "" // Set a default value if phone number is not available
	}

	price, ok := mariner_certificate.Price()
	if !ok {
		price = 1 // Set a default value if phone number is not available
	}

	reg_fee, ok := mariner_certificate.RegFee()
	if !ok {
		reg_fee = 1 // Set a default value if phone number is not available
	}
	updateCertificate, err := prisma.MarinerCertificates.FindUnique(
		db.MarinerCertificates.ID.Equals(id),
	).Update(
		db.MarinerCertificates.Certificate.Link(
			db.Certificates.ID.Equals(mariner_certificate.CertificateID),
		),
		db.MarinerCertificates.Department.Link(
			db.Departments.ID.Equals(mariner_certificate.DepartmentID),
		),
		db.MarinerCertificates.Mariner.Link(
			db.Mariners.ID.Equals(mariner_certificate.MarinerID),
		),
		db.MarinerCertificates.Ship.Link(
			db.Ships.ID.Equals(mariner_certificate.ShipID),
		),
		db.MarinerCertificates.PutinDate.Set(putin_date),
		db.MarinerCertificates.CertID.Set(cert_ID),
		db.MarinerCertificates.IssueDate.Set(issue_date),
		db.MarinerCertificates.ExpireDate.Set(expire_date),
		db.MarinerCertificates.Account.Set(account),
		db.MarinerCertificates.Price.Set(price),
		db.MarinerCertificates.RegFee.Set(reg_fee),
	).Exec(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateCertificate}})
}

func DeleteAMarinerCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	_, err := prisma.MarinerCertificates.FindUnique(db.MarinerCertificates.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Certificate successfully deleted!"}})
}

func GetAllMarinerCertificates(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mariner_certificates, err := prisma.MarinerCertificates.FindMany().With(
		db.MarinerCertificates.Ship.Fetch(),
		db.MarinerCertificates.Department.Fetch(),
		db.MarinerCertificates.Certificate.Fetch(),
		db.MarinerCertificates.Mariner.Fetch(),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": mariner_certificates}})
}
