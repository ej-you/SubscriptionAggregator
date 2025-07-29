package v1

import (
	fiber "github.com/gofiber/fiber/v2"
)

// RegisterSubsEndpoints registers all endpoints for subs entity.
func RegisterSubsEndpoints(router fiber.Router, controller *SubsController) {
	crudlPrefix := router.Group("/subs")

	crudlPrefix.Post("/", controller.Create)
	crudlPrefix.Get("/:id", controller.GetByID)
	crudlPrefix.Patch("/:id", controller.Update)
	crudlPrefix.Delete("/:id", controller.Delete)
	crudlPrefix.Get("/", controller.GetList)

	router.Get("/subs-sum", controller.GetSum)
}
