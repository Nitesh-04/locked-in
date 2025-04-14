package controllers

import (
	"time"
	"github.com/Nitesh-04/locked-in/config"
	"github.com/Nitesh-04/locked-in/models"

	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
)

type StartSessionInput struct {
	ClerkID string    `json:"clerkId"`
	GroupID uuid.UUID `json:"groupId"`
}

type EndSessionInput struct {
	SessionID uuid.UUID `json:"sessionId"`
}

func StartSession(c *fiber.Ctx) error {
	var input StartSessionInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	session := models.StudySession{
		ID:        uuid.New(),
		UserID:   input.ClerkID,
		GroupID:   input.GroupID,
		StartedAt: time.Now(),
	}

	if err := config.DB.Create(&session).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to start session"})
	}


	var user models.User
	result := config.DB.Where("clerk_id = ?",session.UserID).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	user.Status = "studying"
	config.DB.Save(&user)

	var populatedSession models.StudySession
	if err := config.DB.Preload("User").Preload("Group").First(&populatedSession, "id = ?", session.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load session details",
		})
	}

	return c.Status(fiber.StatusOK).JSON(populatedSession)

}

func EndSession(c *fiber.Ctx) error {
	var input EndSessionInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	var session models.StudySession

	if err := config.DB.First(&session, "id = ?", input.SessionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Session not found",
		})
	}

	endTime := time.Now()
	duration := int(endTime.Sub(session.StartedAt).Minutes())

	session.EndedAt = &endTime
	session.Duration = &duration

	if err := config.DB.Save(&session).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to end session",
		})
	}

	var user models.User

	result := config.DB.Where("clerk_id = ?", session.UserID).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	user.Status = "offline"
	config.DB.Save(&user)

	var populatedSession models.StudySession
	if err := config.DB.Preload("User").Preload("Group").First(&populatedSession, "id = ?", session.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load session details",
		})
	}

	return c.Status(fiber.StatusOK).JSON(populatedSession)
}