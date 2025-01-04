package repository

import (
	"database/sql"
	"fmt"
	"valeraninja/noteapp/internal/models"
)

type ItemPostgres struct {
	db *sql.DB
}

func NewItemPostgres(db *sql.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) CreateItem(note models.Note) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s(title, description) values($1, $2) RETURNING id", noteTable)
	row := r.db.QueryRow(query, note.Title, note.Description)
	if err := row.Scan(&id); err != nil{
		return 0, nil
	}
	return id, nil
}

func (r *ItemPostgres) GetAllItems() ([]models.Note, error){
	notes := make([]models.Note, 0)

	query := fmt.Sprintf("select * from %s", noteTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Description, &note.CreatedAt); err != nil{
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *ItemPostgres) GetItemById(id int) (models.Note, error){
	var note models.Note
	query := fmt.Sprintf("select * from %s where id = $1", noteTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&note.ID, &note.Title, &note.Description, &note.CreatedAt); err != nil{
		return models.Note{}, err
	}
	return note, nil
}

func (r *ItemPostgres) UpdateItem(id int, note models.Note) error{
	query := fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2 WHERE id = $3", noteTable, "title", "description")
	result, err := r.db.Exec(query, note.Title, note.Description, note.ID)
	if err != nil{
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		return err
	}

	if rowsAffected == 0{
		return fmt.Errorf("record not found")
	}
	return nil
}

func (r *ItemPostgres) DeleteItem(id int) error{
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", noteTable)

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
	return nil
}