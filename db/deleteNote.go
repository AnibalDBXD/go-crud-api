package db

import (
	"fmt"
	"time"
)

type DeleteNoteParams struct {
	ID int `json:"id"`
}

func DeleteNote(params *DeleteNoteParams) (SuccessResponse, error) {
	dbInstance := getPostgressInstance()
	deleteAt := time.Now()

	_, err := dbInstance.Exec(
		`UPDATE notes SET
			deleted_at = $1
			WHERE id = $2
		`,
		deleteAt,
		params.ID,
	)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return SuccessResponse{}, err
	}

	return SuccessResponse{
		Message: "Note deleted successfully",
	}, nil
}
