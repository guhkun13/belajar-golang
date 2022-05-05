package noteHandler

import (
	"log"

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
		ret := fiber.Map{"status": "error","message" : "no data was found","data": nil}
		log.Println(ret)
	
		return c.Status(404).JSON(ret)
	}
	ret := fiber.Map{"status": "success", "message" : "Data found", "data": notes}
	log.Println(ret)

	return c.JSON(ret)
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note
	
	// read param
	id := c.Params("id")
	
	// search db
	db.Find(&note, "id = ?", id)
	
	if note.ID == uuid.Nil {
		ret := fiber.Map{"status": "error","message" : "no data was found","data": nil}
		return c.Status(404).JSON(ret)
	}
	
	ret := fiber.Map{"status": "success", "message" : "Data found", "data": note}
	return c.JSON(ret)
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB

	note := new(model.Note)

	err := c.BodyParser(note)

	if err != nil {
		ret := fiber.Map{"status": "error","message" : "Review your input","data": err}
		return c.Status(500).JSON(ret)
	}
	
	note.ID = uuid.New()
	err = db.Create(&note).Error

	if err != nil {
		ret := fiber.Map{"status": "error","message" : "Cannot create data","data": err}
		return c.Status(500).JSON(ret)
	}

	ret := fiber.Map{"status": "success", "message" : "Data created", "data": note}
	return c.JSON(ret)	
}
