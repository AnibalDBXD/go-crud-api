package utils

import "github.com/gofiber/fiber/v2"

func SetCorsHeaders(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	return c.Next()
}
