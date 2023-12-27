package main

import (
	"fmt"
	"github.com/aattola/gotest/routes/fetch"
	"github.com/aattola/gotest/routes/kissa"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initOrm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Db kusi")
	}

	return db
}

func main() {
	fmt.Println("kissa")

	app := fiber.New(fiber.Config{})
	db := initOrm()

	kissaGroup := app.Group("kissa")
	kissa.HandleKissa(kissaGroup, *db)

	fetchGroup := app.Group("fetch")
	fetchRouteHandler.HandleFetchRoutes(fetchGroup, *db)

	app.Get("/", func(ctx *fiber.Ctx) error {
		argument := ctx.Query("kissa", "false")

		copyOfArgument := argument

		fmt.Println(copyOfArgument)
		fmt.Println(argument)

		copyOfArgument = "jepa"

		fmt.Println(copyOfArgument)
		fmt.Println(argument)

		return ctx.JSON(&fiber.Map{
			argument: argument,
		})
	})

	app.Get("/map", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"kissa": true,
		})
	})

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
		panic("Kaikki meni rikki!")
	}
}
