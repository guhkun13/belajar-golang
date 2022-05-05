package noteHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/guhkun13/belajar-golang/notes-api/database"
	"github.com/guhkun13/belajar-golang/notes-api/internal/model"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	
	var notes []model.Note
	
	db.Find(&notes)
	
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error","message" : "no data was found","data": nil})
	}
	
	return c.JSON(fiber.Map{"status": "success", "message" : "Data found", "data": notes})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note
	
	// read param
	id := c.Params("id")
	
	// search db
	db.Find(&note, "id = ?", id)
	
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error","message" : "no data was found","data": nil})
	}
	
	
	return c.JSON(fiber.Map{"status": "success", "message" : "Data found", "data": note})
}