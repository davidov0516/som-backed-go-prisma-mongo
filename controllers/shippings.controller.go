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
func CreateShipping(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var shipping db.ShippingsModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&shipping); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a shipping
	port, ok := shipping.Port()
	if !ok {
		port = ""
	}
	b_l, ok := shipping.BL()
	if !ok {
		b_l = ""
	}
	cargo, ok := shipping.Cargo()
	if !ok {
		cargo = ""
	}
	departure_datetime, ok := shipping.DepartureDatetime()
	if !ok {
		var defaultTime time.Time
		departure_datetime = defaultTime
	}
	departure_pass_datetime, ok := shipping.DeparturePassDatetime()
	if !ok {
		var defaultTime time.Time
		departure_pass_datetime = defaultTime
	}
	arrived_datetime, ok := shipping.ArrivedDatetime()
	if !ok {
		var defaultTime time.Time
		arrived_datetime = defaultTime
	}
	arrived_pass_datetime, ok := shipping.ArrivedPassDatetime()
	if !ok {
		var defaultTime time.Time
		arrived_pass_datetime = defaultTime
	}
	shipping_fee, ok := shipping.ShippingFee()
	if !ok {
		shipping_fee = 0
	}
	deposit, ok := shipping.Deposit()
	if !ok {
		deposit = 0
	}
	daily_wages, ok := shipping.DailyWages()
	if !ok {
		daily_wages = 0
	}
	additional_fee, ok := shipping.AdditionalFee()
	if !ok {
		additional_fee = 0
	}
	cost_others_note, ok := shipping.CostOthersNote()
	if !ok {
		cost_others_note = ""
	}
	cost_fees_note, ok := shipping.CostFeesNote()
	if !ok {
		cost_fees_note = ""
	}
	inventory_before_departure, ok := shipping.InventoryBeforeDeparture()
	if !ok {
		inventory_before_departure = 0
	}
	added_fuel, ok := shipping.AddedFuel()
	if !ok {
		added_fuel = 0
	}
	consume_before_departure, ok := shipping.ConsumeBeforeDeparture()
	if !ok {
		consume_before_departure = 0
	}
	inventory_when_arrived, ok := shipping.InventoryWhenArrived()
	if !ok {
		inventory_when_arrived = 0
	}
	fuel_note, ok := shipping.FuelNote()
	if !ok {
		fuel_note = ""
	}
	attachment, ok := shipping.Attachment()
	if !ok {
		attachment = ""
	}
	note, ok := shipping.Note()
	if !ok {
		note = ""
	}

	createdShipping, err := prisma.Shippings.CreateOne(
		db.Shippings.SNumber.Set(shipping.SNumber),
		db.Shippings.Ship.Link(
			db.Ships.ID.Equals(shipping.ShipID),
		),
		db.Shippings.Charterer.Link(
			db.Charterers.ID.Equals(shipping.ChartererID),
		),
		db.Shippings.Port.Set(port),
		db.Shippings.BL.Set(b_l),
		db.Shippings.Cargo.Set(cargo),
		db.Shippings.DepartureDatetime.Set(departure_datetime),
		db.Shippings.DeparturePassDatetime.Set(departure_pass_datetime),
		db.Shippings.ArrivedDatetime.Set(arrived_datetime),
		db.Shippings.ArrivedPassDatetime.Set(arrived_pass_datetime),
		db.Shippings.ShippingFee.Set(shipping_fee),
		db.Shippings.Deposit.Set(deposit),
		db.Shippings.DailyWages.Set(daily_wages),
		db.Shippings.AdditionalFee.Set(additional_fee),
		db.Shippings.CostOthersNote.Set(cost_others_note),
		db.Shippings.CostFeesNote.Set(cost_fees_note),
		db.Shippings.InventoryBeforeDeparture.Set(inventory_before_departure),
		db.Shippings.AddedFuel.Set(added_fuel),
		db.Shippings.ConsumeBeforeDeparture.Set(consume_before_departure),
		db.Shippings.InventoryWhenArrived.Set(inventory_when_arrived),
		db.Shippings.FuelNote.Set(fuel_note),
		db.Shippings.Attachment.Set(attachment),
		db.Shippings.Note.Set(note),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdShipping, "", "  ")
	fmt.Printf("created shipping: %s\n", result)

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAShipping(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	shipping, err := prisma.Shippings.FindUnique(
		db.Shippings.ID.Equals(id),
	).With(
		db.Shippings.Ship.Fetch(),
		db.Shippings.Charterer.Fetch(),
		db.Shippings.CrewFees.Fetch().With(
			db.CrewFees.Mariner.Fetch().With(
				db.Mariners.Job.Fetch(),
			),
		),
		db.Shippings.OtherCosts.Fetch(),
		db.Shippings.ShippingAreas.Fetch().With(
			db.ShippingAreas.WaterArea.Fetch(),
			db.ShippingAreas.LoadUnloads.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": shipping}})
}

func EditAShipping(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")

	var shipping db.ShippingsModel
	defer cancel()

	//parse the request body
	if err := c.BodyParser(&shipping); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// create a shipping
	port, ok := shipping.Port()
	if !ok {
		port = ""
	}
	b_l, ok := shipping.BL()
	if !ok {
		b_l = ""
	}
	cargo, ok := shipping.Cargo()
	if !ok {
		cargo = ""
	}
	departure_datetime, ok := shipping.DepartureDatetime()
	if !ok {
		var defaultTime time.Time
		departure_datetime = defaultTime
	}
	departure_pass_datetime, ok := shipping.DeparturePassDatetime()
	if !ok {
		var defaultTime time.Time
		departure_pass_datetime = defaultTime
	}
	arrived_datetime, ok := shipping.ArrivedDatetime()
	if !ok {
		var defaultTime time.Time
		arrived_datetime = defaultTime
	}
	arrived_pass_datetime, ok := shipping.ArrivedPassDatetime()
	if !ok {
		var defaultTime time.Time
		arrived_pass_datetime = defaultTime
	}
	shipping_fee, ok := shipping.ShippingFee()
	if !ok {
		shipping_fee = 0
	}
	deposit, ok := shipping.Deposit()
	if !ok {
		deposit = 0
	}
	daily_wages, ok := shipping.DailyWages()
	if !ok {
		daily_wages = 0
	}
	additional_fee, ok := shipping.AdditionalFee()
	if !ok {
		additional_fee = 0
	}
	cost_others_note, ok := shipping.CostOthersNote()
	if !ok {
		cost_others_note = ""
	}
	cost_fees_note, ok := shipping.CostFeesNote()
	if !ok {
		cost_fees_note = ""
	}
	inventory_before_departure, ok := shipping.InventoryBeforeDeparture()
	if !ok {
		inventory_before_departure = 0
	}
	added_fuel, ok := shipping.AddedFuel()
	if !ok {
		added_fuel = 0
	}
	consume_before_departure, ok := shipping.ConsumeBeforeDeparture()
	if !ok {
		consume_before_departure = 0
	}
	inventory_when_arrived, ok := shipping.InventoryWhenArrived()
	if !ok {
		inventory_when_arrived = 0
	}
	fuel_note, ok := shipping.FuelNote()
	if !ok {
		fuel_note = ""
	}
	attachment, ok := shipping.Attachment()
	if !ok {
		attachment = ""
	}
	note, ok := shipping.Note()
	if !ok {
		note = ""
	}

	updateShipping, err := prisma.Shippings.FindUnique(
		db.Shippings.ID.Equals(id),
	).Update(
		db.Shippings.SNumber.Set(shipping.SNumber),
		db.Shippings.Ship.Link(
			db.Ships.ID.Equals(shipping.ShipID),
		),
		db.Shippings.Charterer.Link(
			db.Charterers.ID.Equals(shipping.ChartererID),
		),
		db.Shippings.Port.Set(port),
		db.Shippings.BL.Set(b_l),
		db.Shippings.Cargo.Set(cargo),
		db.Shippings.DepartureDatetime.Set(departure_datetime),
		db.Shippings.DeparturePassDatetime.Set(departure_pass_datetime),
		db.Shippings.ArrivedDatetime.Set(arrived_datetime),
		db.Shippings.ArrivedPassDatetime.Set(arrived_pass_datetime),
		db.Shippings.ShippingFee.Set(shipping_fee),
		db.Shippings.Deposit.Set(deposit),
		db.Shippings.DailyWages.Set(daily_wages),
		db.Shippings.AdditionalFee.Set(additional_fee),
		db.Shippings.CostOthersNote.Set(cost_others_note),
		db.Shippings.CostFeesNote.Set(cost_fees_note),
		db.Shippings.InventoryBeforeDeparture.Set(inventory_before_departure),
		db.Shippings.AddedFuel.Set(added_fuel),
		db.Shippings.ConsumeBeforeDeparture.Set(consume_before_departure),
		db.Shippings.InventoryWhenArrived.Set(inventory_when_arrived),
		db.Shippings.FuelNote.Set(fuel_note),
		db.Shippings.Attachment.Set(attachment),
		db.Shippings.Note.Set(note),
	).Exec(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateShipping}})
}

func DeleteAShipping(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Params("id")
	defer cancel()

	_, err := prisma.Shippings.FindUnique(db.Shippings.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Shipping successfully deleted!"}})
}

func GetAllShippings(c *fiber.Ctx) error {
	prisma := configs.PrismaClient

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	shippings, err := prisma.Shippings.FindMany().With(
		db.Shippings.Ship.Fetch(),
		db.Shippings.Charterer.Fetch(),
		db.Shippings.CrewFees.Fetch().With(
			db.CrewFees.Mariner.Fetch(),
		),
		db.Shippings.OtherCosts.Fetch(),
		db.Shippings.ShippingAreas.Fetch().With(
			db.ShippingAreas.WaterArea.Fetch(),
			db.ShippingAreas.LoadUnloads.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": shippings}})
}
