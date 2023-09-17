package controllers

import "github.com/gofiber/fiber/v2"

func InsertMoney(c *fiber.Ctx) error {
	done := make(chan bool)
	go func(){
		
	}()
	<-done
	return c.JSON(map[string]interface{}{
		"status": true,
		"message": "",
		"data": nil,
	})
}