package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nitesh-04/locked-in/controllers"
)

func MembershipRoutes(app *fiber.App) {

	membership := app.Group("/group")
	membership.Post("/join", controllers.JoinGroup)

}