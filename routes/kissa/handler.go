package kissa

import (
	"github.com/gofiber/fiber/v2"
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

		db.Create(&uusiKissa)

		return ctx.JSON(&uusiKissa)
	})

	return nil
}
