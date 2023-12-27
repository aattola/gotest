package kissa

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type KissaModel struct {
	gorm.Model
	Nimi  string
	Hinta uint64
	ID    string
}

func HandleKissa(r fiber.Router, db gorm.DB) error {

	db.AutoMigrate(&KissaModel{})

	r.Get("/map", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"kissa": true,
		})
	})

	r.Get("/luo", func(ctx *fiber.Ctx) error {

		uusiKissa := &KissaModel{
			Nimi:  "Kass",
			Hinta: 22,
			ID:    "123",
		}

		tx := db.Create(&uusiKissa)

		if tx.Error != nil {

			fmt.Println(tx.Error)
			sqliteErr := tx.Error.(sqlite3.Error)
			switch sqliteErr.Code {
			case 19:
				return ctx.JSON(&fiber.Map{
					"error": "a table constraint (NOT NULL, UNIQUE, etc.) was violated during the operation (INSERT, etc.)",
				})
			}

			return ctx.JSON(&fiber.Map{
				"error": tx.Error,
				"err":   tx.Error.Error(),
			})
		}

		return ctx.JSON(&uusiKissa)
	})

	r.Get("/luo/autoid", func(ctx *fiber.Ctx) error {

		uusiKissa := &KissaModel{
			Nimi:  "Kass",
			Hinta: 22,
		}

		db.Create(&uusiKissa)

		return ctx.JSON(&uusiKissa)
	})

	r.Get("/hae/:id", func(ctx *fiber.Ctx) error {

		id := ctx.Params("id")

		var kissa KissaModel

		tx := db.First(&kissa, id)

		if tx.Error != nil {
			return ctx.JSON(&fiber.Map{
				"error": tx.Error,
				"err":   tx.Error.Error(),
			})
		}

		return ctx.JSON(&kissa)
	})

	return nil
}
