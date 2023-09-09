package handlers

import (
	"fmt"

	"github.com/AnibalDBXD/go-crud-api/db"
	"github.com/AnibalDBXD/go-crud-api/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleCreateNote(c *fiber.Ctx) error {
	params, err := utils.GetBodyParams[db.CreateNoteParams](c)

	if err != nil {
		return utils.ResponseWithError(c, 400, err.Error())
	}

	note, err := db.CreateNote(&db.CreateNoteParams{
		Title:       params.Title,
		Description: params.Description,
	})

	if err != nil {
		fmt.Println("HandleCreateNote error: ", err)
		return utils.ResponseWithError(c, 500, err.Error())
	}
	fmt.Println("HandleCreateNote success: ", note)
	return utils.ResponseWithJSON(c, 200, note)
}
