package handlers

import (
	"fmt"
	"strconv"

	"github.com/AnibalDBXD/go-crud-api/db"
	"github.com/AnibalDBXD/go-crud-api/utils"
	"github.com/gofiber/fiber/v2"
)

// note/:id
func HandleGetNoteById(c *fiber.Ctx) error {
	id := c.Params("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil {
		return utils.ResponseWithError(c, 400, "Invalid id")
	}

	note, queryErr := db.GetNoteById(&db.GetNoteByIdParams{
		Id: idNumber,
	})

	if queryErr.Error() == "Note not found" {
		return utils.ResponseWithError(c, 404, "Note not found")
	}

	if queryErr != nil {
		fmt.Println("HandleGetNoteById error: ", queryErr)
		return utils.ResponseWithError(c, 500, queryErr.Error())
	}
	fmt.Println("HandleGetNoteById success: ", note)
	return utils.ResponseWithJSON(c, 200, note)
}
