package services

import (
	"valeraninja/noteapp/internal/models"
	"valeraninja/noteapp/internal/repository"
)

type NoteItemService struct {
	repo repository.NoteItem
}

func NewNoteItemService(repo repository.NoteItem) *NoteItemService {
	return &NoteItemService{repo: repo}
}

func (s *NoteItemService) CreateItem(note models.Note) (int, error) {
	return s.repo.CreateItem(note)
}

func (s *NoteItemService) GetAllItems() ([]models.Note, error) {
	return s.repo.GetAllItems()
}

func (s *NoteItemService) GetItemById(id int) (models.Note, error) {
	return s.repo.GetItemById(id)
}

func (s *NoteItemService) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}

func (s *NoteItemService) UpdateItem(id int, note models.Note) error {
	return s.repo.UpdateItem(id, note)
}
