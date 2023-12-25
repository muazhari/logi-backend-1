package controllers

import (
	"logi-backend-1/src/inners/use_cases/managements"

	"github.com/gofiber/fiber/v2"
)

type LogController struct {
	router        fiber.Router
	logManagement *managements.LogManagement
}

func NewLogController(router fiber.Router, logManagement *managements.LogManagement) *LogController {
	logController := &LogController{
		router:        router,
		logManagement: logManagement,
	}
	return logController
}

func (logController *LogController) RegisterRoutes() {
	logController.router.Post("/logs", logController.CreateOne)
	logController.router.Get("/logs/:id", logController.ReadOneById)
	logController.router.Get("/logs", logController.ReadMany)
	logController.router.Put("/logs/:id", logController.UpdateOneById)
	logController.router.Delete("/logs/:id", logController.DeleteOneById)
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
