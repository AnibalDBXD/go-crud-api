package handlers

import (
	"github.com/AnibalDBXD/go-crud-api/utils"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return utils.ResponseWithJSON(c, 200, map[string]string{
		"message": "Health check OK",
	})
}
