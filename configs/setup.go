package configs

import (
	"log"
	"som-backend/prisma/db"
)

var PrismaClient *db.PrismaClient

func ConnectPrismaClient() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	PrismaClient = client
}

func DisconnectPrismaClient() {
	if err := PrismaClient.Prisma.Disconnect(); err != nil {
		panic(err)
	}
}
