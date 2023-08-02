package main

import (
	"example/aggrt_channels/handlers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app:=fiber.New()

	app.Get("/calculate/:num1/:num2",handlers.FindAggrt)

	err:=app.Listen(":3000")
	if err!=nil{
		fmt.Println("Error in starting SERVER:",err)
	}
}