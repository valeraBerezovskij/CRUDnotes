package services

import (
	"valeraninja/noteapp/internal/models"
	"valeraninja/noteapp/internal/repository"
)

type NoteItem interface {
	CreateItem(note models.Note) (int, error)
	GetAllItems() ([]models.Note, error)
	GetItemById(id int) (models.Note, error)
	DeleteItem(id int) error
	UpdateItem(id int, note models.Note) error
}

type Service struct {
	NoteItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NoteItem: NewNoteItemService(repos.NoteItem),
	}
}
