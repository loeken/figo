package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/loeken/figo/pkg/handler"
	"gorm.io/gorm"
)

// Content content
func Content(app *fiber.App, db *gorm.DB) {
	h := &handler.ContentHandler{
		DB: db,
	}
	r := app.Group("/api/v1/content")
	r.Get("/", h.Index)
	r.Get("/:id", h.Show)
	r.Post("/", h.Store)
	r.Put("/:id", h.Update)
	r.Delete("/:id", h.Destroy)
}
