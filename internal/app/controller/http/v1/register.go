package v1

import (
	fiber "github.com/gofiber/fiber/v2"
)

// RegisterSubsEndpoints registers all endpoints for subs entity.
func RegisterSubsEndpoints(router fiber.Router, controller *SubsController) {
	router.Post("/", controller.Create)
	router.Get("/:id", controller.GetByID)
	router.Patch("/:id", controller.Update)
	router.Delete("/:id", controller.Delete)
	router.Get("/", controller.GetAll)

	router.Get("/sum", controller.GetSum)
}
