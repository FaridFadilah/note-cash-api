package routes

import (
	"fmt"
	"main/constants"
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){
	msg := fmt.Sprintf("Server running in port : %v", constants.PORT)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"status" : true,
			"message": msg,
			"data": nil,
		})
	})

	app.Post("/signin", controllers.Authentication)	
	app.Post("/register", controllers.Register)	
}