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
func CreateCertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var certificate db.CertificatesModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&certificate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a certificate
	note, ok := certificate.Note()
	if !ok {
		note = "" // Set a default value if phone number is not available
	}

	createdCertificate, err := prisma.Certificates.CreateOne(
		db.Certificates.Name.Set(certificate.Name),
		db.Certificates.AgencyName.Set(certificate.AgencyName),
		db.Certificates.Type.Set(certificate.Type),
		db.Certificates.Note.Set(note),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdCertificate, "", "  ")
	fmt.Printf("created certificate: %s\n", result)

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})

	// res := map[string]interface{}{
	// 	"name":  "John Doe",
	// 	"email": "john@example.com",
	// 	"age":   30,
	// }

	// return c.Status(http.StatusOK).JSON(res)
}

func GetACertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	certificate, err := prisma.Certificates.FindUnique(db.Certificates.ID.Equals(id)).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": certificate}})
}

func EditACertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")

	var certificate db.CertificatesModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&certificate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	note, ok := certificate.Note()
	if !ok {
		note = "" // Set a default value if phone number is not available
	}

	updateCertificate, err := prisma.Certificates.FindMany(
		db.Certificates.ID.Equals(id),
	).Update(
		db.Certificates.Name.Set(certificate.Name),
		db.Certificates.AgencyName.Set(certificate.AgencyName),
		db.Certificates.Type.Set(certificate.Type),
		db.Certificates.Note.Set(note),
	).Exec(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateCertificate}})
}

func DeleteACertificate(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	_, err := prisma.Certificates.FindUnique(db.Certificates.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Certificate successfully deleted!"}})
}

func GetAllCertificates(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	certificates, err := prisma.Certificates.FindMany().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": certificates}})
}
