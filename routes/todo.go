package routes

import "github.com/gofiber/fiber/v2"
import "github.com/nguyenhuuluan434/gofiber/controllers"

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
