package main

import (
	"context"
	"fmt"
	"log"
	"main/constants"
	"main/models"
	"main/routes"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){ 
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	models.ConnectDB()

	routes.Routes()
	
	signals := []os.Signal{
		os.Interrupt,
		os.Kill,
	}

	route := []http.HandlerFunc{}

	srv := &http.Server{
		Addr: ":3000",
	}
  
	fmt.Println(constants.IMG)
	fmt.Printf("Server is running in port :%v", 3000)
	fmt.Printf("The numbers of Endpoint: %v", )

	go func(){
		if err := srv.ListenAndServe(); err != nil{
			log.Println(err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
  signal.Notify(done, signals...)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
	 log.Fatal(err)
	}

	fmt.Println("Server has been stopped")
}