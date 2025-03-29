package controllers

import (
	"github.com/Nitesh-04/locked-in/config"
	"github.com/Nitesh-04/locked-in/models"

	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	req := new(models.User)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	var existingUser models.User
	if err := config.DB.Where("clerk_id = ?", req.ClerkID).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User already exists",
			"user":    existingUser,
		})
	}

	newUser := models.User{
		ID : uuid.New(),
		ClerkID: req.ClerkID,
		Name : req.Name,
		Email : req.Email,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    newUser,
	})
}