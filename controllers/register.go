package controllers

import (
	"main/helpers"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	db := models.ConnectDB()

	user := new(models.User)
	
	if err := c.BodyParser(user); err != nil{
		panic(err)
	}
	
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	
	user.Password = helpers.Encrypt(user.Password, "password")
	user.UUID = helpers.Encrypt(user.Password, "uuid")
	db.Create(&user)

	return c.JSON(map[string]interface{}{
		"status" 	: false,
		"message"	: "Success",
		"data"		: user,
	})
}