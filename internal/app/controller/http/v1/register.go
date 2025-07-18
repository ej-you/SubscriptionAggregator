package v1

import (
	fiber "github.com/gofiber/fiber/v2"
)

// RegisterSubsEndpoints registers all endpoints for subs entity.
func RegisterSubsEndpoints(router fiber.Router, controller *SubsController) {
	router.Post("/subs", controller.Create)
	router.Get("/subs/:id", controller.GetByID)
	router.Patch("/subs/:id", controller.Update)
	router.Delete("/subs/:id", controller.Delete)
	router.Get("/subs", controller.GetAll)

	router.Get("/subs-sum", controller.GetSum)
}
