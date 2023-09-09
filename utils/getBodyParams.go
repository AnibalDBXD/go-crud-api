package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetBodyParams[K any](c *fiber.Ctx) (K, error) {
	body := c.Body()
	var params K
	err := json.Unmarshal(body, &params)
	if err != nil {
		return params, err
	}

	return params, nil
}
