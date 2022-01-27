package router

import (
	"github.com/OShuaib/Human-resource-management/handler"
	"github.com/gofiber/fiber"
)

func SetupRoutes (app *fiber.App) { 
     
	api := app.Group("/api")
    // routes
    api.Get("/employees", handler.GetEmployee)
    api.Post("/employee", handler.CreateEmployee)
    api.Put("/employee/:id", handler.UpdateEmployee)
    api.Delete("/employee/:id", handler.DeleteEmployee)
}