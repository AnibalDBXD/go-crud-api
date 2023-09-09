package handlers

import (
	"fmt"
	"strconv"

	"github.com/AnibalDBXD/go-crud-api/db"
	"github.com/AnibalDBXD/go-crud-api/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleGetNotes(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return utils.ResponseWithError(c, 400, "Invalid page")
	}
	limit := c.Query("limit", "10")
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		return utils.ResponseWithError(c, 400, "Invalid limit")
	}
	title := c.Query("title", "")
	description := c.Query("description", "")

	createdAtFrom := c.Query("created_at_from", "")
	createdAtTo := c.Query("created_at_to", "")

	updateAtFrom := c.Query("updated_at_from", "")
	updateAtTo := c.Query("updated_at_to", "")

	response, queryErr := db.GetNotes(&db.GetNotesParams{
		Page:          pageNumber,
		Limit:         limitNumber,
		Title:         title,
		Description:   description,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updateAtFrom,
		UpdatedAtTo:   updateAtTo,
	})

	if queryErr != nil {
		fmt.Println("HandleGetNotes error: ", queryErr)
		return utils.ResponseWithError(c, 500, queryErr.Error())
	}
	fmt.Println("HandleGetNotes success")
	return utils.ResponseWithJSON(c, 200, response)
}
