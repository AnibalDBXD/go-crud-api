package db

import "fmt"

type GetNoteByIdParams struct {
	Id int
}

func GetNoteById(params *GetNoteByIdParams) (Note, error) {
	dbInstance := getPostgressInstance()
	response, err := dbInstance.Query(
		`SELECT
			id, title, description, created_at, updated_at
		FROM notes
		WHERE id = $1 and deleted_at is NULL
		`,
		params.Id,
	)

	if err != nil {
		fmt.Printf("Error: %v", err)
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
	// check if note is empty
	if note.ID == 0 {
		return Note{}, fmt.Errorf("Note not found")
	}

	fmt.Println("GetNoteById success: ", note)

	return note, nil
}
