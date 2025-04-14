package controllers

import (
	"github.com/Nitesh-04/locked-in/config"
	"github.com/Nitesh-04/locked-in/models"

	"github.com/gofiber/fiber/v2"
)

func JoinGroup(c *fiber.Ctx) error {

	type RequestBody struct {
		Name	string `json:"name"`
		ClerkId string `json:"clerkId"`
	}

	var req RequestBody

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	var group models.Group

	if err := config.DB.Where("name = ?", req.Name).First(&group).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Group not found",
		})
	}

	var membership models.GroupMembership

	if err := config.DB.Where("user_id = ? AND group_id = ?", req.ClerkId, group.ID).First(&membership).Error; err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User already a member of the group",
			"group":   group,
		})
	}

	membership = models.GroupMembership{
		UserID:  req.ClerkId,
		GroupID: group.ID,
	}

	if err := config.DB.Create(&membership).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to join group",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User joined the group successfully",
		"group":   group,
	})
	
}



