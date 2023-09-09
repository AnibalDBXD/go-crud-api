package db

import "fmt"

type GetNotesParams struct {
	Page          int
	Limit         int
	Title         string
	Description   string
	CreatedAtFrom string
	CreatedAtTo   string
	UpdatedAtFrom string
	UpdatedAtTo   string
}

type GetNotesResult struct {
	Notes []Note `json:"notes"`
	Total int    `json:"total"`
}

func GetNotes(params *GetNotesParams) (GetNotesResult, error) {
	dbInstance := getPostgressInstance()

	pageSize := params.Limit

	// calculate the offset based on the page number and page size
	offset := (params.Page - 1) * pageSize

	// query to get the notes for the current page
	query := `
		SELECT id, title, description, created_at, updated_at
		FROM notes
	`

	args := make([]interface{}, 0)
	args = append(args, pageSize)
	args = append(args, offset)

	count := 3

	// if the title is not empty, add the where clause to the query
	if params.Title != "" {
		query += fmt.Sprintf(" WHERE title LIKE $%v", count)
		args = append(args, fmt.Sprintf("%%%v%%", params.Title))
		count++
	}

	// if the description is not empty, add the where clause to the query
	if params.Description != "" {
		query += fmt.Sprintf(" WHERE description LIKE $%v ", count)
		args = append(args, fmt.Sprintf("%%%v%%", params.Description))
		count++
	}

	// if the created_at_from is not empty, add the where clause to the query
	if params.CreatedAtFrom != "" {
		query += fmt.Sprintf(" WHERE created_at >= $%v ", count)
		args = append(args, params.CreatedAtFrom)
		count++
	}

	// if the created_at_to is not empty, add the where clause to the query
	if params.CreatedAtTo != "" {
		query += fmt.Sprintf(" WHERE created_at <= $%v ", count)
		args = append(args, params.CreatedAtTo)
		count++
	}

	// if the updated_at_from is not empty, add the where clause to the query
	if params.UpdatedAtFrom != "" {
		query += fmt.Sprintf(" WHERE updated_at >= $%v ", count)
		args = append(args, params.UpdatedAtFrom)
		count++
	}

	// if the updated_at_to is not empty, add the where clause to the query
	if params.UpdatedAtTo != "" {
		query += fmt.Sprintf(" WHERE updated_at <= $%v ", count)
		args = append(args, params.UpdatedAtTo)
		count++
	}

	query = fmt.Sprintf(`
		%v
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`, query)

	rows, err := dbInstance.Query(query,
		args...,
	)

	if err != nil {
		fmt.Printf("Error on get query: %v", err)
		return GetNotesResult{}, err
	}
	defer rows.Close()

	// query to get the total number of notes
	var total int
	err = dbInstance.QueryRow(
		`
					SELECT COUNT(*)
					FROM notes
			`,
	).Scan(&total)

	if err != nil {
		fmt.Printf("Error on get count: %v", err)
		return GetNotesResult{}, err
	}

	// iterate over the rows and add them to the notes slice
	var notes []Note
	for rows.Next() {
		var note Note
		err = rows.Scan(&note.ID, &note.Title, &note.Description, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return GetNotesResult{}, err
		}
		notes = append(notes, note)
	}

	// create the result object and return it
	result := GetNotesResult{
		Notes: notes,
		Total: total,
	}
	return result, nil
}
