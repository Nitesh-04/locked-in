package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nitesh-04/locked-in/controllers"
)

func GroupRoutes(app *fiber.App) {

	groupGroup := app.Group("/group")
	groupGroup.Post("/", controllers.CreateGroup)
	groupGroup.Get("/:groupID", controllers.GetGroupInfo)
}