package main

import (
	"fmt"
	"log"

	"github.com/OShuaib/Human-resource-management/Db"
	"github.com/OShuaib/Human-resource-management/router"
	"github.com/gofiber/fiber"
)



func main(){
if err:= database.Connect(); err != nil {
	log.Fatal(err)
}
 
	app := fiber.New()

	router.SetupRoutes(app)

	fmt.Println(app.Listen(":300"))
}