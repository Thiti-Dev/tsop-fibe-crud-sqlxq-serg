package main

import (
	"fmt"

	"github.com/Thiti-Dev/go-fiber-101/database"
	"github.com/Thiti-Dev/go-fiber-101/models"
	"github.com/Thiti-Dev/go-fiber-101/utils"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main(){
	app := fiber.New()
	app.Use(middleware.Logger("${blue} ${time} ${red} ${latency} ${cyan} ${method} ${path} ${yellow} ${status} ${green} ${ip} ${ua}${resetColor}"))

	app.Post("api/insert", func(c *fiber.Ctx){
		//req := 

		req := new(models.UserModelReq)
		err := utils.DoesReqBodyCanParsed(req,c)
		if err != nil{
			return
		}

		qr := fmt.Sprintf(`INSERT INTO users ("UserID","FirstName","LastName","Age", "Status","CreatedAt")` + 
			`VALUES(uuid_generate_v4(), '%v', '%v', %d, %d, current_timestamp)` , req.FirstName, req.LastName, req.Age, req.Status)
		fmt.Println(qr)
		err = database.DoCrudOperationWithTargetQuery(qr)
		if err != nil{
			utils.MakeErrorRespondIfAnyError(err,c)
			return
		}


		c.Status(201).JSON(&fiber.Map{
			"status": "success",
			"data": req,
		})

	})


	app.Listen(3000)
}

