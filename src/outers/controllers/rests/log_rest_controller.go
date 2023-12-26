package rests

import (
	"github.com/muazhari/logi-backend-1/src/inners/use_cases/managements"

	"github.com/gofiber/fiber/v2"
)

type LogController struct {
	Router        fiber.Router
	LogManagement *managements.LogManagement
}

func NewLogController(router fiber.Router, logManagement *managements.LogManagement) *LogController {
	logController := &LogController{
		Router:        router,
		LogManagement: logManagement,
	}
	return logController
}

func (logController *LogController) RegisterRoutes() {
	logController.Router.Post("/logs", logController.CreateOne)
	logController.Router.Get("/logs/:id", logController.ReadOneById)
	logController.Router.Get("/logs", logController.ReadMany)
	logController.Router.Put("/logs/:id", logController.UpdateOneById)
	logController.Router.Delete("/logs/:id", logController.DeleteOneById)
}

func (logController *LogController) CreateOne(ctx *fiber.Ctx) error {
	return nil
}

func (logController *LogController) ReadOneById(ctx *fiber.Ctx) error {
	return nil
}

func (logController *LogController) ReadMany(ctx *fiber.Ctx) error {
	return nil
}

func (logController *LogController) UpdateOneById(ctx *fiber.Ctx) error {
	return nil
}

func (logController *LogController) DeleteOneById(ctx *fiber.Ctx) error {
	return nil
}
