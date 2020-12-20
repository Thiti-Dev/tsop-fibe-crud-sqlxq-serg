package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main(){
	app := fiber.New()
	app.Use(middleware.Logger("${blue} ${time} ${red} ${latency} ${cyan} ${method} ${path} ${yellow} ${status} ${green} ${ip} ${ua}${resetColor}"))

	


	app.Listen(3000)
}

