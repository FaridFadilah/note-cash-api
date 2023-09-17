package controllers

import (
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	done := make(chan bool)
	go func(){
		db := models.ConnectDB()
		h := c.GetReqHeaders()
		db.Delete(&models.AuthToken{}, "\"token\" = ?", h["Authorization"])
		done <- true	
	}()
	<-done
	return c.JSON(map[string]interface{}{
		"status": true,
		"message": "",
		"data": nil,
	})
}