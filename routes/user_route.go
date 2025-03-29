package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nitesh-04/locked-in/controllers"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Post("/", controllers.CreateUser)
}