package db

import (
	"fmt"
	"time"
)

type CreateNoteParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateNote(params *CreateNoteParams) (Note, error) {
	dbInstance := getPostgressInstance()

	createdAt := time.Now()
	updatedAt := time.Now()

	_, err := dbInstance.Exec(
		"INSERT INTO notes (title, description, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		params.Title,
		params.Description,
		createdAt,
		updatedAt,
	)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return Note{}, err
	}

	return Note{
		Title:       params.Title,
		Description: params.Description,
		CreatedAt:   createdAt.String(),
		UpdatedAt:   updatedAt.String(),
	}, nil
}
