package controllers

import (
	"main/models"
	"net/http"
)

func CreateUpdateWallet(w http.ResponseWriter, r *http.Request) error {
	db := models.ConnectDB()
	wallet := new(models.Wallet)
	
	return c.JSON(map[string]interface{}{
		"status": false,
		"message": "",
		"data": nil,
	})
}