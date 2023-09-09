package main

import (
	"fmt"

	"github.com/AnibalDBXD/go-crud-api/handlers"
	"github.com/AnibalDBXD/go-crud-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	envs, error := utils.GetEnvs()
	if error != nil {
		fmt.Println(error)
		return
	}

	app := fiber.New()
	app.Use(utils.SetCorsHeaders)
	app.Get("/healthcheck", handlers.HealthCheck)
	app.Post("/note", handlers.HandleCreateNote)
	app.Get("/notes", handlers.HandleGetNotes)
	app.Get("/note/:id", handlers.HandleGetNoteById)
	app.Put("/note/:id", handlers.HandleUpdateNote)
	app.Delete("/note/:id", handlers.HandleDeleteNote)
	app.Listen(":" + envs.Port)
}
