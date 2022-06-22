package control

import "github.com/gofiber/fiber/v2"

func Register(httpPath fiber.Router) {
	router := httpPath.Group("/control")

	router.Post("/load", Upload)
}
