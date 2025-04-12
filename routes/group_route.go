package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nitesh-04/locked-in/controllers"
)

func GroupRoutes(app *fiber.App) {

	group := app.Group("/group")
	group.Post("/", controllers.CreateGroup)
	group.Get("/:groupID", controllers.GetGroupInfo)
	group.Get("/:groupID/users", controllers.GetGroupUsers)
	group.Delete("/:groupID", controllers.DeleteGroup)
}