package utils

import (
	"encoding/json"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type TestParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestGetBodyParams(t *testing.T) {
	// create a test fiber context with a request body
	params := TestParams{Name: "John Doe", Email: "john.doe@example.com"}
	body, _ := json.Marshal(params)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Set("Content-Type", "application/json")
	ctx.Request().SetBody(body)

	// call the GetBodyParams function with the test context
	result, err := GetBodyParams[TestParams](ctx)

	// assert that the result and error are as expected
	assert.Nil(t, err)
	assert.Equal(t, params, result)
}
