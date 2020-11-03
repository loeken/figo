package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/loeken/figo/pkg/entity"
	"github.com/loeken/figo/pkg/validation"
	"gorm.io/gorm"
)

// ContentHandler type
type ContentHandler struct {
	DB *gorm.DB
}

// Index to list all content
func (h ContentHandler) Index(ctx *fiber.Ctx) error {
	var content []entity.Content
	h.DB.Find(&content)
	return ctx.JSON(fiber.Map{"data": content})
}

// Show a content
func (h ContentHandler) Show(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your content"}})
	}

	var contentDB entity.Content
	contentDB.ID = uint(id)
	h.DB.First(&contentDB)

	return ctx.JSON(fiber.Map{"data": contentDB})
}

// Store a new content
func (h ContentHandler) Store(ctx *fiber.Ctx) error {
	content := new(entity.Content)

	if err := ctx.BodyParser(content); err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
	}

	if validate := validation.CreateOrUpdateContent(*content); len(validate) > 0 {
		return ctx.Status(422).JSON(fiber.Map{"errors": validate})
	}

	h.DB.Create(&content)
	return ctx.JSON(fiber.Map{"message": "Content successfully added"})
}

// Update an content
func (h ContentHandler) Update(ctx *fiber.Ctx) error {
	content := new(entity.Content)

	if err := ctx.BodyParser(content); err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
	}

	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your content"}})
	}

	if validate := validation.CreateOrUpdateContent(*content); len(validate) > 0 {
		return ctx.Status(422).JSON(fiber.Map{"errors": validate})
	}

	var contentDB entity.Content
	h.DB.First(&contentDB, id)
	h.DB.Model(&contentDB).Updates(map[string]interface{}{
		"title":  content.Title,
		"body": content.Body,
	})

	return ctx.JSON(fiber.Map{"message": "Content successfully updated"})
}

// Destroy an content
func (h ContentHandler) Destroy(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your content"}})
	}

	var contentDB entity.Content
	contentDB.ID = uint(id)
	h.DB.Delete(&contentDB)

	return ctx.JSON(fiber.Map{"message": "Content successfully removed"})
}

