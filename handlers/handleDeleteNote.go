package handlers

import (
	"fmt"
	"strconv"

	"github.com/AnibalDBXD/go-crud-api/db"
	"github.com/AnibalDBXD/go-crud-api/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleDeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil {
		return utils.ResponseWithError(c, 400, "Invalid id")
	}

	if err != nil {
		return utils.ResponseWithError(c, 400, err.Error())
	}

	note, err := db.DeleteNote(&db.DeleteNoteParams{
		ID: idNumber,
	})

	if err != nil {
		fmt.Println("HandleDeleteNote error: ", err)
		return utils.ResponseWithError(c, 500, err.Error())
	}
	fmt.Println("HandleDeleteNote success: ", note)
	return utils.ResponseWithJSON(c, 200, note)
}
