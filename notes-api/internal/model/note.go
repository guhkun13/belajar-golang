package model 

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Note struct {
	gorm.Model
	ID 				uuid.UUID `gorm:"type=uuid"`
	Title 		string 
	Subtitle 	string
	Text 			string

}
