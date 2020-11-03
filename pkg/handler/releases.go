package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/loeken/figo/pkg/entity"
	"github.com/loeken/figo/pkg/validation"
	"gorm.io/gorm"
)

// ReleaseHandler type
type ReleaseHandler struct {
	DB *gorm.DB
}

// Index to list all releases
func (h ReleaseHandler) Index(ctx *fiber.Ctx) error {
	var releases []entity.Release
	h.DB.Find(&releases)
	return ctx.JSON(fiber.Map{"data": releases})
}

// Show a release
func (h ReleaseHandler) Show(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your release"}})
	}

	var releaseDB entity.Release
	releaseDB.ID = uint(id)
	h.DB.First(&releaseDB)

	return ctx.JSON(fiber.Map{"data": releaseDB})
}

// Store a new release
func (h ReleaseHandler) Store(ctx *fiber.Ctx) error {
	release := new(entity.Release)

	if err := ctx.BodyParser(release); err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
	}

	if validate := validation.CreateOrUpdateRelease(*release); len(validate) > 0 {
		return ctx.Status(422).JSON(fiber.Map{"errors": validate})
	}

	h.DB.Create(&release)
	return ctx.JSON(fiber.Map{"message": "Release successfully added"})
}

// Update an release
func (h ReleaseHandler) Update(ctx *fiber.Ctx) error {
	release := new(entity.Release)

	if err := ctx.BodyParser(release); err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
	}

	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your release"}})
	}

	if validate := validation.CreateOrUpdateRelease(*release); len(validate) > 0 {
		return ctx.Status(422).JSON(fiber.Map{"errors": validate})
	}

	var releaseDB entity.Release
	h.DB.First(&releaseDB, id)
	h.DB.Model(&releaseDB).Updates(map[string]interface{}{
		"title":  release.Title,
		"artist": release.Artist,
		"label":  release.Label,
	})

	return ctx.JSON(fiber.Map{"message": "Release successfully updated"})
}

// Destroy an release
func (h ReleaseHandler) Destroy(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your release"}})
	}

	var releaseDB entity.Release
	releaseDB.ID = uint(id)
	h.DB.Delete(&releaseDB)

	return ctx.JSON(fiber.Map{"message": "Release successfully removed"})
}

// Upload an attachment
func (h ReleaseHandler) Upload(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your release"}})
	}

	file, err := ctx.FormFile("attachment")

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})
	}

	ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

	var releaseDB entity.Release
	h.DB.First(&releaseDB, id)
	h.DB.Model(&releaseDB).Update("attachment", file.Filename)

	return ctx.JSON(fiber.Map{"message": "Attachment uploaded successfully"})
}
