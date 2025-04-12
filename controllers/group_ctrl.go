package controllers

import (
	"github.com/Nitesh-04/locked-in/config"
	"github.com/Nitesh-04/locked-in/models"

	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
)

func CreateGroup(c *fiber.Ctx) error {
	type RequestBody struct {
		Name    string `json:"name"`
		ClerkID string `json:"clerkId"`
	}
	var req RequestBody

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	var existingGroup models.Group
	if err := config.DB.Where("name = ?", req.Name).First(&existingGroup).Error; err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Group with name already exists",
			"group":   existingGroup,
		})
	}

	newGroup := models.Group{
		ID:   uuid.New(),
		Name: req.Name,
	}

	if err := config.DB.Create(&newGroup).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create group",
		})
	}

	membership := models.GroupMembership{
		ID:       uuid.New(),
		UserID:   req.ClerkID,
		GroupID:  newGroup.ID,
	}

	if err := config.DB.Create(&membership).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Group created but failed to add user to group",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Group created successfully",
		"group":   newGroup,
	})
}


func GetGroupInfo(c *fiber.Ctx) error {
	groupId := c.Params("groupID")

	var group models.Group
	result := config.DB.Where("id = ?", groupId).First(&group)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Group not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"group": group,
	})
}

func GetGroupUsers(c *fiber.Ctx) error {
	groupId := c.Params("groupID")

	var group models.Group
	result := config.DB.Where("id = ?", groupId).First(&group)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Group not found",
		})
	}

	var users []models.User
	if err := config.DB.Model(&models.GroupMembership{}).
		Where("group_id = ?", group.ID).
		Joins("JOIN users ON users.id = group_memberships.user_id").
		Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve group members",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}


func DeleteGroup(c *fiber.Ctx) error {
	groupId := c.Params("groupID")

	var group models.Group
	result := config.DB.Where("id = ?", groupId).First(&group)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Group not found",
		})
	}

	if err := config.DB.Delete(&group).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete group",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Group deleted successfully",
	})
}