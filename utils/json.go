package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ResponseWithJSON(c *fiber.Ctx, code int, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling payload: %v", err)
	}

	c.Set("Content-Type", "application/json")
	c.Status(code)
	c.Send(data)

	return nil
}

func ResponseWithError(c *fiber.Ctx, code int, message string) error {
	return ResponseWithJSON(c, code, map[string]string{"error": message})
}
