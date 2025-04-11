package controllers

import (
	"github.com/Nitesh-04/locked-in/config"
	"github.com/Nitesh-04/locked-in/models"

	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
)

func JoinGroup(c *fiber.Ctx) error {

	groupIdParam := c.Params("groupID")
	groupId,err := uuid.Parse(groupIdParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	type RequesBody struct {
		ClerkId string `json:"clerk_id"`
	}

	var reqBody RequesBody

	if err := c.BodyParser(&reqBody); err != nil || reqBody.ClerkId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or invalid clerkID",
		})
	}

	var group models.Group

	if err := config.DB.Where("id = ?", groupId).First(&group).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Group not found",
		})
	}

	var existing models.GroupMembership

	err = config.DB.Where("group_id = ? AND clerk_id = ?", groupId, reqBody.ClerkId).First(&existing).Error

	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already a member of the group",
		})
	}

	var count int64

	config.DB.Model(&models.GroupMembership{}).Where("group_id = ?", groupId).Count(&count)

	if count >= 10 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Group is full",
		})
	}

	membership := models.GroupMembership{
		UserID: reqBody.ClerkId,
		GroupID: groupId,
	}

	if err := config.DB.Create(&membership).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to join group",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Joined group successfully",
	})
}



