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

}
