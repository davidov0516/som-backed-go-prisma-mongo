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
func CreateCharterer(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var charterer db.CharterersModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&charterer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a charterer
	phoneNumber, ok := charterer.PhoneNumber()
	if !ok {
		phoneNumber = "" // Set a default value if phone number is not available
	}
	emailNumber, ok := charterer.EmailNumber()
	if !ok {
		emailNumber = "" // Set a default value if phone number is not available
	}
	note, ok := charterer.Note()
	if !ok {
		note = "" // Set a default value if phone number is not available
	}

	createdCharterer, err := prisma.Charterers.CreateOne(
		db.Charterers.CompanyName.Set(charterer.CompanyName),
		db.Charterers.Nation.Set(charterer.Nation),
		db.Charterers.PhoneNumber.Set(phoneNumber),
		db.Charterers.EmailNumber.Set(emailNumber),
		db.Charterers.Note.Set(note),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdCharterer, "", "  ")
	fmt.Printf("created charterer: %s\n", result)

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})

	// res := map[string]interface{}{
	// 	"name":  "John Doe",
	// 	"email": "john@example.com",
	// 	"age":   30,
	// }

	// return c.Status(http.StatusOK).JSON(res)
}

func GetACharterer(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	charterer, err := prisma.Charterers.FindUnique(db.Charterers.ID.Equals(id)).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": charterer}})
}

func EditACharterer(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")

	var charterer db.CharterersModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&charterer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a charterer
	phoneNumber, ok := charterer.PhoneNumber()
	if !ok {
		phoneNumber = "" // Set a default value if phone number is not available
	}
	emailNumber, ok := charterer.EmailNumber()
	if !ok {
		emailNumber = "" // Set a default value if phone number is not available
	}
	note, ok := charterer.Note()
	if !ok {
		note = "" // Set a default value if phone number is not available
	}

	updateCharterer, err := prisma.Charterers.FindMany(
		db.Charterers.ID.Equals(id),
	).Update(
		db.Charterers.CompanyName.Set(charterer.CompanyName),
		db.Charterers.Nation.Set(charterer.Nation),
		db.Charterers.PhoneNumber.Set(phoneNumber),
		db.Charterers.EmailNumber.Set(emailNumber),
		db.Charterers.Note.Set(note),
	).Exec(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateCharterer}})
}

func DeleteACharterer(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	_, err := prisma.Charterers.FindUnique(db.Charterers.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Charterer successfully deleted!"}})
}

func GetAllCharterers(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	charterers, err := prisma.Charterers.FindMany().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": charterers}})
}
