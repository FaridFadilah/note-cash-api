package main

import (
	"fmt"
	"main/constants"
	"main/models"
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func main(){ 
	models.ConnectDB()
	app := fiber.New()
	
	routes.Routes(app)

	app.Listen(fmt.Sprintf(":%v", constants.PORT))
}