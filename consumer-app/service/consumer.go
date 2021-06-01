package service

import (
	"github.com/pranaySinghDev/gRPCConnectionLoadTest/consumer-app/utils"

	"github.com/gofiber/fiber/v2"
)

func Base(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func GetUser(c *fiber.Ctx) error {
	response := utils.GetUser()
	return c.Status(200).JSON(response)
}

func AddUser(c *fiber.Ctx) error {
	response := utils.AddUser()
	return c.Status(201).JSON(response)
}
