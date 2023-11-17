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
func CreateShip(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var ship db.ShipsModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&ship); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a ship
	ship_type, ok := ship.Type()
	if !ok {
		ship_type = ""
	}
	year_of_build, ok := ship.YearOfBuild()
	if !ok {
		year_of_build = 0
	}
	flag, ok := ship.Flag()
	if !ok {
		flag = ""
	}
	homeport, ok := ship.Homeport()
	if !ok {
		homeport = ""
	}
	reg_number, ok := ship.RegNumber()
	if !ok {
		reg_number = ""
	}
	callsign, ok := ship.Callsign()
	if !ok {
		callsign = ""
	}
	IMO_Number, ok := ship.IMONumber()
	if !ok {
		IMO_Number = ""
	}
	gross_tonnage, ok := ship.GrossTonnage()
	if !ok {
		gross_tonnage = 0
	}
	net_tonnage, ok := ship.NetTonnage()
	if !ok {
		net_tonnage = 0
	}
	deadweight, ok := ship.Deadweight()
	if !ok {
		deadweight = 0
	}
	length, ok := ship.Length()
	if !ok {
		length = 0
	}
	beam, ok := ship.Beam()
	if !ok {
		beam = 0
	}
	depth, ok := ship.Depth()
	if !ok {
		depth = 0
	}
	draught, ok := ship.Draught()
	if !ok {
		draught = 0
	}
	note, ok := ship.Note()
	if !ok {
		note = ""
	}
	// photo, ok := ship.Photo()
	// if !ok {
	// 	photo = []string
	// }

	createdShip, err := prisma.Ships.CreateOne(
		db.Ships.Name.Set(ship.Name),
		db.Ships.RemovedDate.Set(ship.RemovedDate),
		db.Ships.RegisteredDate.Set(ship.RegisteredDate),
		db.Ships.IsRemoved.Set(ship.IsRemoved),
		db.Ships.Type.Set(ship_type),
		db.Ships.YearOfBuild.Set(year_of_build),
		db.Ships.Flag.Set(flag),
		db.Ships.Homeport.Set(homeport),
		db.Ships.RegNumber.Set(reg_number),
		db.Ships.Callsign.Set(callsign),
		db.Ships.IMONumber.Set(IMO_Number),
		db.Ships.GrossTonnage.Set(gross_tonnage),
		db.Ships.NetTonnage.Set(net_tonnage),
		db.Ships.Deadweight.Set(deadweight),
		db.Ships.Length.Set(length),
		db.Ships.Beam.Set(beam),
		db.Ships.Depth.Set(depth),
		db.Ships.Draught.Set(draught),
		db.Ships.Note.Set(note),
		db.Ships.Photo.Set(ship.Photo),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdShip, "", "  ")
	fmt.Printf("created ship: %s\n", result)

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAShip(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	ship, err := prisma.Ships.FindUnique(db.Ships.ID.Equals(id)).With(
		db.Ships.ShipCertificates.Fetch().With(
			db.ShipCertificates.Certificate.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": ship}})
}

func EditAShip(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")

	var ship db.ShipsModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&ship); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a ship
	ship_type, ok := ship.Type()
	if !ok {
		ship_type = ""
	}
	year_of_build, ok := ship.YearOfBuild()
	if !ok {
		year_of_build = 0
	}
	flag, ok := ship.Flag()
	if !ok {
		flag = ""
	}
	homeport, ok := ship.Homeport()
	if !ok {
		homeport = ""
	}
	reg_number, ok := ship.RegNumber()
	if !ok {
		reg_number = ""
	}
	callsign, ok := ship.Callsign()
	if !ok {
		callsign = ""
	}
	IMO_Number, ok := ship.IMONumber()
	if !ok {
		IMO_Number = ""
	}
	gross_tonnage, ok := ship.GrossTonnage()
	if !ok {
		gross_tonnage = 0
	}
	net_tonnage, ok := ship.NetTonnage()
	if !ok {
		net_tonnage = 0
	}
	deadweight, ok := ship.Deadweight()
	if !ok {
		deadweight = 0
	}
	length, ok := ship.Length()
	if !ok {
		length = 0
	}
	beam, ok := ship.Beam()
	if !ok {
		beam = 0
	}
	depth, ok := ship.Depth()
	if !ok {
		depth = 0
	}
	draught, ok := ship.Draught()
	if !ok {
		draught = 0
	}
	note, ok := ship.Note()
	if !ok {
		note = ""
	}
	// photo, ok := ship.Photo()
	// if !ok {
	// 	photo = []string
	// }

	updateShip, err := prisma.Ships.FindMany(
		db.Ships.ID.Equals(id),
	).Update(
		db.Ships.Name.Set(ship.Name),
		db.Ships.RemovedDate.Set(ship.RemovedDate),
		db.Ships.RegisteredDate.Set(ship.RegisteredDate),
		db.Ships.IsRemoved.Set(ship.IsRemoved),
		db.Ships.Type.Set(ship_type),
		db.Ships.YearOfBuild.Set(year_of_build),
		db.Ships.Flag.Set(flag),
		db.Ships.Homeport.Set(homeport),
		db.Ships.RegNumber.Set(reg_number),
		db.Ships.Callsign.Set(callsign),
		db.Ships.IMONumber.Set(IMO_Number),
		db.Ships.GrossTonnage.Set(gross_tonnage),
		db.Ships.NetTonnage.Set(net_tonnage),
		db.Ships.Deadweight.Set(deadweight),
		db.Ships.Length.Set(length),
		db.Ships.Beam.Set(beam),
		db.Ships.Depth.Set(depth),
		db.Ships.Draught.Set(draught),
		db.Ships.Note.Set(note),
		db.Ships.Photo.Set(ship.Photo),
	).Exec(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateShip}})
}

func DeleteAShip(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	_, err := prisma.Ships.FindUnique(db.Ships.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Ship successfully deleted!"}})
}

func GetAllShips(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ships, err := prisma.Ships.FindMany().With(
		db.Ships.ShipCertificates.Fetch().With(
			db.ShipCertificates.Certificate.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": ships}})
}
