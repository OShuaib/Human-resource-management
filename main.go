package main

import (
	
	"log"
	"github.com/OShuaib/Human-resource-management/handler"
	"github.com/OShuaib/Human-resource-management/Db"
	"github.com/gofiber/fiber/v2"
	
)







func main(){
if err:= database.Connect(); err != nil {
	log.Fatal(err)
}
 
	app := fiber.New()

	app.Get("/employee", handler.GetEmployee())


	app.Post("/employee", handler.CreateEmployee())


	app.Put("/employee/:id",handler.UpdateEmployee())


	app.Delete("/employee/:id", handler.DeleteEmployee())

	app.Listen(":300")
}