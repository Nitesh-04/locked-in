package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nitesh-04/locked-in/controllers"
)

func GroupRoutes(app *fiber.App) {

	groupGroup := app.Group("/group")
	groupGroup.Post("/", controllers.CreateGroup)
	groupGroup.Get("/:groupName", controllers.GetGroupInfo)
	groupGroup.Get("/:groupName/users", controllers.GetGroupUsers)
	groupGroup.Delete("/:groupName", controllers.DeleteGroup)
}