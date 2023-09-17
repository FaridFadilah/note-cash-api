package routes

import (
	"net/http"
)

func Routes(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
	})
	// msg := fmt.Sprintf("Server running in port : %v", constants.PORT)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(map[string]interface{}{
	// 		"status" : true,
	// 		"message": msg,
	// 		"data": nil,
	// 	})
	// })

	// // Authentication
	// app.Post("/v1/signin", controllers.Authentication)	
	// app.Post("/v1/register", controllers.Register)	
	// app.Post("/v1/logout", controllers.Logout)

	// // wallet
	// app.Post("/v1/wallet/create", controllers.CreateUpdateWallet)	
	// app.Post("/v1/wallet/update", controllers.CreateUpdateWallet)	
}