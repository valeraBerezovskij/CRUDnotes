package repository

import (
	"database/sql"
	"valeraninja/noteapp/internal/models"
)

type NoteItem interface {
	CreateItem(note models.Note) (int, error)
	GetAllItems() ([]models.Note, error)
	GetItemById(id int) (models.Note, error)
	DeleteItem(id int) error
	UpdateItem(id int, note models.Note) error
}

type Repository struct {
	NoteItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NoteItem: NewItemPostgres(db),
	}
}
