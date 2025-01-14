package repository

import (
	"database/sql"
	"fmt"
	"time"
	"valeraninja/noteapp/internal/models"
	"valeraninja/noteapp/pkg/database"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type ItemPostgres struct {
	db    *sql.DB
	cache *cache.Cache
}

func NewItemPostgres(db *sql.DB) *ItemPostgres {
	return &ItemPostgres{
		db:    db,
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (r *ItemPostgres) CreateItem(note models.Note) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s(title, description) values($1, $2) RETURNING id", database.NoteTable)
	row := r.db.QueryRow(query, note.Title, note.Description)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	r.cache.Delete("GetAllItems")
	return id, nil
}

func (r *ItemPostgres) GetAllItems() ([]models.Note, error) {
	if cached, found := r.cache.Get("GetAllItems"); found {
		logrus.Infoln("cache get all items")
		return cached.([]models.Note), nil
	}

	notes := make([]models.Note, 0)
	query := fmt.Sprintf("select * from %s", database.NoteTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Description, &note.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	r.cache.Set("GetAllItems", notes, cache.DefaultExpiration)
	return notes, nil
}

func (r *ItemPostgres) GetItemById(id int) (models.Note, error) {
	cacheKey := fmt.Sprintf("GetItemById:%d", id)
	if cached, found := r.cache.Get(cacheKey); found {
		logrus.Infoln("cache GetItemById")
		return cached.(models.Note), nil
	}

	var note models.Note
	query := fmt.Sprintf("select * from %s where id = $1", database.NoteTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&note.ID, &note.Title, &note.Description, &note.CreatedAt); err != nil {
		return models.Note{}, err
	}

	r.cache.Set(cacheKey, note, cache.DefaultExpiration)

	return note, nil
}

func (r *ItemPostgres) UpdateItem(id int, note models.Note) error {
	query := fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2 WHERE id = $3", database.NoteTable, "title", "description")
	result, err := r.db.Exec(query, note.Title, note.Description, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("record not found")
	}

	r.cache.Delete(fmt.Sprintf("GetItemById:%d", id))
	r.cache.Delete("GetAllItems")
	return nil
}

func (r *ItemPostgres) DeleteItem(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", database.NoteTable)

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("record not found")
	}

	r.cache.Delete(fmt.Sprintf("GetItemById:%d", id))
	r.cache.Delete("GetAllItems")

	return nil
}
