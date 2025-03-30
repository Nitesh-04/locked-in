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

func GetUserInfo(c *fiber.Ctx) error {
	clerkId := c.Params("clerkID")

	var user models.User
	result := config.DB.Where("clerk_id = ?", clerkId).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func UpdateUserStatus(c *fiber.Ctx) error {

	clerkID := c.Params("clerkId")

	var request struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	validStatuses := map[string]bool{
		"studying":  true,
		"on_break":  true,
		"offline":   true,
	}

	if !validStatuses[request.Status] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status",
		})
	}

	var user models.User
	result := config.DB.Where("clerk_id = ?", clerkID).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	user.Status = request.Status
	config.DB.Save(&user)

	return c.JSON(user)
}

func GetUserGroups(c *fiber.Ctx) error {
	
	clerkID := c.Params("clerkId")

	var user models.User
	result := config.DB.Where("clerk_id = ?", clerkID).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	var groupMembers []models.GroupMembership

	config.DB.Where("user_id = ?", user.ID).Find(&groupMembers)
	
	var groupIDs []uuid.UUID

	for _, gm := range groupMembers {
		groupIDs = append(groupIDs, gm.GroupID)
	}

	var groups []models.Group
	config.DB.Where("id IN ?", groupIDs).Find(&groups)
	
	return c.JSON(groups)
}