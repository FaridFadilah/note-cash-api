package controllers

import (
	"fmt"
	"log"
	"main/helpers"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Authentication(c *fiber.Ctx) error {
	db := models.ConnectDB()

	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil{
		panic(err)
	}
	
	var userDb models.User
	db.Table("users").Where("\"email\"=?", user.Email).Select("id", "uuid", "fullName", "email", "password").First(&userDb)

	if userDb.ID == 0{
		log.Println("Email & password invalid")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"status"	: false,
			"message"	:	"Email & password invalid", 
			"Data"		: nil,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))

	if err != nil{
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"status": false,
			"message": "Email & password invalid",
			"data": nil,
		})
	}
	
	var userToken models.AuthToken
	userToken.UserId = int64(userDb.ID)
	userToken.Token = helpers.Encrypt(fmt.Sprintf("%v", time.Now().Nanosecond()), "token")
	db.Table("auth_tokens").Create(&userToken)

	if userDb.ID == 0 || err != nil{
		log.Println("Username & password is wrong")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"status": false,
			"message": "Username & Password is wrong",
			"data": nil, 
		})
	}

	return c.JSON(map[string]interface{}{
		"status": true,
		"message": "Success",
		"data": map[string]interface{}{
			"name": userDb.FullName,
			"email": userDb.Email,
			"kind": userDb.Role,
			"token": userToken.Token,
		},
	})
}