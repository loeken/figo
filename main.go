package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/loeken/figo/pkg/connector"
	"github.com/loeken/figo/pkg/route"
)

func main() {
	app := fiber.New()
	// switch to sqlite
	db := connector.ConnectSqlite()
	// switch to mysql
	// db := connector.ConnectMysql()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Figo API",
		})
	})

	route.Releases(app, db)
	route.Content(app, db)
	app.Listen(":3000")

}
