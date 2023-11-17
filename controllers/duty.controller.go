package controllers

import (
	"context"
	"net/http"
	"som-backend/configs"
	"som-backend/prisma/db"
	"som-backend/responses"
	"time"

	"github.com/gofiber/fiber/v2"
)


func GetADuty(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	duty, err := prisma.Duties.FindUnique(db.Duties.ID.Equals(id)).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": duty}})
}


// func DeleteADuty(c *fiber.Ctx) error {
// 	prisma := configs.PrismaClient

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	id := c.Params("id")
// 	defer cancel()

// 	_, err := prisma.WaterAreas.FindUnique(db.WaterAreas.ID.Equals(id)).Delete().Exec(ctx)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "duty successfully deleted!"}})
// }

func GetAllDuties(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	duties, err := prisma.Duties.FindMany().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": duties}})
}
