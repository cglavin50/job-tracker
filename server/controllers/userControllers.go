package controllers

import (
	"github.com/cglavin50/job-tracker/server/models"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var body models.Job
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to request body")
	}

	return nil
}
