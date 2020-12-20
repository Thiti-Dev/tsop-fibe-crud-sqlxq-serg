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

	app.Post("api/update", func(c *fiber.Ctx){
		//req := 

		req := new(models.UserModelUpdate)
		err := utils.DoesReqBodyCanParsed(req,c)
		if err != nil{
			return
		}

		qr := fmt.Sprintf(`UPDATE users ` +
			`SET "FirstName"='%v', "LastName"='%v', "Age"=%d, "Status"=%d, "CreatedAt"=current_timestamp ` +
			`WHERE "UserID"='%v'`, req.FirstName, req.LastName, req.Age, req.Status, req.UserID)
		fmt.Println(qr)
		err = database.DoCrudOperationWithTargetQuery(qr)
		if err != nil{
			utils.MakeErrorRespondIfAnyError(err,c)
			return
		}


		c.Status(200).JSON(&fiber.Map{
			"message": "success",
			"data": req,
		})

	})


	app.Get("api/select", func(c *fiber.Ctx){
		db, err := database.OpenPgSQL()
		if err != nil{
			fmt.Println(err)
		}
		qr := `select * from users`
		selects,err := db.Query(qr)
		_ = selects

		var model []models.UserModel = make([]models.UserModel, 0)
		var modelDb models.UserModelDb

		for selects.Next() {

			// This scan function -> aligned weirdly because it has to be ordered as the same as the table alignment in order of postgresql
			err := selects.Scan(&modelDb.Age,&modelDb.Status,&modelDb.UserID,&modelDb.FirstName,&modelDb.LastName,&modelDb.CreatedAt)
			if err != nil {
				fmt.Println(err)
			}
			modelSet := models.UserModel{modelDb.UserID.String,
				modelDb.FirstName.String, modelDb.LastName.String, int(modelDb.Age.Int32), int(modelDb.Status.Int32), modelDb.CreatedAt.String}
			
			model = append(model, modelSet)
		}

		c.Status(200).JSON(&fiber.Map{
			"status": "success",
			"data": model,
		})
	})

	app.Delete("api/delete/:id", func(c *fiber.Ctx){
		id := c.Params("id")
		qr := fmt.Sprintf(`Delete from users where "UserID"='%v'`, id)
		
		fmt.Println(qr)
		err := database.DoCrudOperationWithTargetQuery(qr)
		if err != nil{
			utils.MakeErrorRespondIfAnyError(err,c)
			return
		}
		c.Status(200).JSON(&fiber.Map{
			"status": "success",
		})
	})


	app.Listen(3000)
}

