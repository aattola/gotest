package fetchRouteHandler

import (
	"fmt"
	"github.com/go-zoox/core-utils/array"
	"github.com/go-zoox/fetch"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleFetchRoutes(r fiber.Router, db gorm.DB) error {
	r.Get("/catfact", func(ctx *fiber.Ctx) error {

		res, err := fetch.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}

		type CatFact struct {
			Fact   string
			Length int
		}

		var catFact CatFact
		if err := res.UnmarshalJSON(&catFact); err != nil {
			return err
		}

		fmt.Println(catFact.Fact)

		return ctx.JSON(&catFact)
	})

	r.Get("/filter", func(ctx *fiber.Ctx) error {

		type KissaObject struct {
			Kissa bool   `json:"kissa"`
			Name  string `json:"name"`
		}

		arr := []KissaObject{KissaObject{Kissa: true, Name: "kissa"}, KissaObject{Kissa: false, Name: "koira"}}

		newarr := array.Filter(arr, func(value KissaObject, index int) bool {
			return value.Kissa
		})

		newarr2 := append(newarr, KissaObject{Kissa: false, Name: "capybara"})

		fmt.Println(newarr2)

		return ctx.JSON(&newarr2)

	})

	return nil
}
