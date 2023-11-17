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
func CreateMariner(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var mariner db.MarinersModel

	defer cancel()

	//parse the request body
	if err := c.BodyParser(&mariner); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a mariner
	daily_fee, ok := mariner.DailyFee()
	if !ok {
		daily_fee = 0 // Set a default value if phone number is not available
	}
	platoon, ok := mariner.Platoon()
	if !ok {
		platoon = ""
	}
	previous_affiliation, ok := mariner.PreviousAffiliation()
	if !ok {
		previous_affiliation = ""
	}
	place_born, ok := mariner.PlaceBorn()
	if !ok {
		place_born = ""
	}
	place_residence, ok := mariner.PlaceResidence()
	if !ok {
		platoon = ""
	}
	code, ok := mariner.Code()
	if !ok {
		code = ""
	}
	mobile_phone, ok := mariner.MobilePhone()
	if !ok {
		mobile_phone = ""
	}
	home_phone, ok := mariner.HomePhone()
	if !ok {
		home_phone = ""
	}
	graduated_from, ok := mariner.GraduatedFrom()
	if !ok {
		graduated_from = ""
	}
	graduated_date, ok := mariner.GraduatedDate()
	if !ok {
		var defaultTime time.Time
		graduated_date = defaultTime
	}

	qualification_grade, ok := mariner.QualificationGrade()
	if !ok {
		qualification_grade = ""
	}
	boarded_years, ok := mariner.BoardedYears()
	if !ok {
		boarded_years = 0
	}
	photo, ok := mariner.Photo()
	if !ok {
		photo = ""
	}
	note, ok := mariner.Note()
	if !ok {
		note = ""
	}

	createdMariner, err := prisma.Mariners.CreateOne(
		db.Mariners.Name.Set(mariner.Name),
		db.Mariners.Birthday.Set(mariner.Birthday),
		db.Mariners.Ship.Link(
			db.Ships.ID.Equals(mariner.ShipID),
		),
		db.Mariners.Duty.Link(
			db.Duties.ID.Equals(mariner.DutyID),
		),
		db.Mariners.Job.Link(
			db.Duties.ID.Equals(mariner.JobID),
		),
		db.Mariners.RegisteredDate.Set(mariner.RegisteredDate),
		db.Mariners.RetiredDate.Set(mariner.RetiredDate),
		db.Mariners.IsRetired.Set(mariner.IsRetired),
		db.Mariners.DailyFee.Set(daily_fee),
		db.Mariners.Platoon.Set(platoon),
		db.Mariners.PreviousAffiliation.Set(previous_affiliation),
		db.Mariners.PlaceBorn.Set(place_born),
		db.Mariners.PlaceResidence.Set(place_residence),
		db.Mariners.Code.Set(code),
		db.Mariners.MobilePhone.Set(mobile_phone),
		db.Mariners.HomePhone.Set(home_phone),
		db.Mariners.GraduatedFrom.Set(graduated_from),
		db.Mariners.GraduatedDate.Set(graduated_date),
		db.Mariners.QualificationGrade.Set(qualification_grade),
		db.Mariners.BoardedYears.Set(boarded_years),
		db.Mariners.Photo.Set(photo),
		db.Mariners.Note.Set(note),
	).Exec(ctx)

	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdMariner, "", "  ")
	fmt.Printf("created mariner: %s\n", result)

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAMariner(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	mariner, err := prisma.Mariners.FindUnique(db.Mariners.ID.Equals(id)).With(
		db.Mariners.Ship.Fetch(),
		db.Mariners.Duty.Fetch(),
		db.Mariners.Job.Fetch(),
		db.Mariners.Certificates.Fetch(),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": mariner}})
}

func EditAMariner(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")

	var mariner db.MarinersModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&mariner); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a mariner
	// create a mariner
	daily_fee, ok := mariner.DailyFee()
	if !ok {
		daily_fee = 0 // Set a default value if phone number is not available
	}
	platoon, ok := mariner.Platoon()
	if !ok {
		platoon = ""
	}
	previous_affiliation, ok := mariner.PreviousAffiliation()
	if !ok {
		previous_affiliation = ""
	}
	place_born, ok := mariner.PlaceBorn()
	if !ok {
		place_born = ""
	}
	place_residence, ok := mariner.PlaceResidence()
	if !ok {
		platoon = ""
	}
	code, ok := mariner.Code()
	if !ok {
		code = ""
	}
	mobile_phone, ok := mariner.MobilePhone()
	if !ok {
		mobile_phone = ""
	}
	home_phone, ok := mariner.HomePhone()
	if !ok {
		home_phone = ""
	}
	graduated_from, ok := mariner.GraduatedFrom()
	if !ok {
		graduated_from = ""
	}
	graduated_date, ok := mariner.GraduatedDate()
	if !ok {
		var defaultTime time.Time
		graduated_date = defaultTime
	}

	qualification_grade, ok := mariner.QualificationGrade()
	if !ok {
		qualification_grade = ""
	}
	boarded_years, ok := mariner.BoardedYears()
	if !ok {
		boarded_years = 0
	}
	photo, ok := mariner.Photo()
	if !ok {
		photo = ""
	}
	note, ok := mariner.Note()
	if !ok {
		note = ""
	}

	updateMariner, err := prisma.Mariners.FindUnique(
		db.Mariners.ID.Equals(id),
	).Update(
		db.Mariners.Name.Set(mariner.Name),
		db.Mariners.Birthday.Set(mariner.Birthday),
		db.Mariners.Ship.Link(
			db.Ships.ID.Equals(mariner.ShipID),
		),
		db.Mariners.Duty.Link(
			db.Duties.ID.Equals(mariner.DutyID),
		),
		db.Mariners.Job.Link(
			db.Duties.ID.Equals(mariner.JobID),
		),
		db.Mariners.RegisteredDate.Set(mariner.RegisteredDate),
		db.Mariners.RetiredDate.Set(mariner.RetiredDate),
		db.Mariners.IsRetired.Set(mariner.IsRetired),
		db.Mariners.DailyFee.Set(daily_fee),
		db.Mariners.Platoon.Set(platoon),
		db.Mariners.PreviousAffiliation.Set(previous_affiliation),
		db.Mariners.PlaceBorn.Set(place_born),
		db.Mariners.PlaceResidence.Set(place_residence),
		db.Mariners.Code.Set(code),
		db.Mariners.MobilePhone.Set(mobile_phone),
		db.Mariners.HomePhone.Set(home_phone),
		db.Mariners.GraduatedFrom.Set(graduated_from),
		db.Mariners.GraduatedDate.Set(graduated_date),
		db.Mariners.QualificationGrade.Set(qualification_grade),
		db.Mariners.BoardedYears.Set(boarded_years),
		db.Mariners.Photo.Set(photo),
		db.Mariners.Note.Set(note),
	).Exec(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateMariner}})
}

func DeleteAMariner(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	_, err := prisma.Mariners.FindUnique(db.Mariners.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Mariner successfully deleted!"}})
}

func GetAllMariners(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mariners, err := prisma.Mariners.FindMany().With(
		db.Mariners.Ship.Fetch(),
		db.Mariners.Duty.Fetch(),
		db.Mariners.Job.Fetch(),
		db.Mariners.MarinerCertificates.Fetch().With(
			db.MarinerCertificates.Certificate.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": mariners}})
}
