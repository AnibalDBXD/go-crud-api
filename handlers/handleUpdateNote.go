package handlers

import (
	"fmt"
	"strconv"

	"github.com/AnibalDBXD/go-crud-api/db"
	"github.com/AnibalDBXD/go-crud-api/utils"
	"github.com/gofiber/fiber/v2"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
func HandleUpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil {
		return utils.ResponseWithError(c, 400, "Invalid id")
	}

	params, err := utils.GetBodyParams[db.UpdateNoteParams](c)

	if err != nil {
		return utils.ResponseWithError(c, 400, err.Error())
	}
	// Print type
	fmt.Print("Type: ", typeof(params.Description), len(params.Description))
	note, err := db.UpdateNote(&db.UpdateNoteParams{
		Title:       params.Title,
		Description: params.Description,
		ID:          idNumber,
	})

	if err != nil {
		fmt.Println("HandleUpdateNote error: ", err)
		return utils.ResponseWithError(c, 500, err.Error())
	}
	fmt.Println("HandleUpdateNote success: ", note)
	return utils.ResponseWithJSON(c, 200, note)
}
