package db

import (
	"fmt"
	"time"
)

type UpdateNoteParams struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateNote(params *UpdateNoteParams) (Note, error) {
	count := 1
	dbInstance := getPostgressInstance()

	updatedAt := time.Now()

	query := "UPDATE notes SET"

	args := make([]interface{}, 0)

	if params.Title != "" {
		query += fmt.Sprintf(" title = $%v,", count)
		args = append(args, params.Title)
		count++
	}

	if params.Description != "" {
		query += fmt.Sprintf(" description = $%v,", count)
		args = append(args, params.Description)
		count++
	}

	if params.Title == "" && params.Description == "" {
		return Note{}, fmt.Errorf("No fields to update")
	}

	query += fmt.Sprintf(`
		updated_at = $%v
		WHERE id = $%v
		RETURNING id, title, description, created_at, updated_at`,
		count,
		count+1,
	)
	args = append(args, updatedAt)
	args = append(args, params.ID)

	response, err := dbInstance.Query(
		query,
		args...,
	)

	if err != nil {
		fmt.Printf("Error on query: %v", err)
		return Note{}, err
	}
	defer response.Close()

	var note Note

	for response.Next() {
		err := response.Scan(
			&note.ID,
			&note.Title,
			&note.Description,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			fmt.Printf("Error on scan: %v", err)
			return Note{}, err
		}
	}

	return note, nil
}
