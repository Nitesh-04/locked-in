package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nitesh-04/locked-in/controllers"
)

func StudySessionRoutes(app *fiber.App) {
	sessionGroup := app.Group("/session")
	sessionGroup.Post("/", controllers.StartSession)
	sessionGroup.Patch("/", controllers.EndSession)
}