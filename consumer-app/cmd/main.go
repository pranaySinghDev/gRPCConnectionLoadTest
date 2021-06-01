package main

import (
	"github.com/pranaySinghDev/gRPCConnectionLoadTest/consumer-app/service"
	"github.com/pranaySinghDev/gRPCConnectionLoadTest/consumer-app/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	utils.InitUserRPC()
	app.Get("/", service.Base)
	app.Get("/user", service.GetUser)
	app.Post("/user", service.AddUser)
	app.Listen(":3000")
}
