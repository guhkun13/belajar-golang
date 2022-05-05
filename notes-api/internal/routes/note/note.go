package noteRoutes

import (
	"github.com/gofiber/fiber/v2"

	noteHandler "github.com/guhkun13/belajar-golang/notes-api/internal/handlers/note"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("note")

	note.Get("/", noteHandler.GetNotes)
	note.Get("/:id", noteHandler.GetNote)
	note.Post("/", noteHandler.CreateNote)
	// note.Put("/:id", noteHandler.UpdateNote)
	// note.Delete("/:id", noteHandler.DeleteNote)
}