package utils

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type TestResponse struct {
	Message string `json:"message"`
}

func TestResponseWithJSON(t *testing.T) {
	// create a test fiber context
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	// create a test payload
	payload := TestResponse{Message: "Hello, World!"}

	// call the ResponseWithJSON function with the test context and payload
	err := ResponseWithJSON(ctx, http.StatusOK, payload)

	// assert that the response is as expected
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, ctx.Response().StatusCode())
	contentType := ctx.Response().Header.Peek("Content-Type")
	assert.Equal(t, "application/json", string(contentType))

	// decode the response body and assert that it is equal to the test payload
	var result TestResponse
	err = json.Unmarshal(ctx.Response().Body(), &result)
	assert.Nil(t, err)
	assert.Equal(t, payload, result)
}

func TestResponseWithError(t *testing.T) {
	// create a test fiber context
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	// call the ResponseWithError function with the test context and error message
	err := ResponseWithError(ctx, http.StatusBadRequest, "Bad Request")

	// assert that the response is as expected
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, ctx.Response().StatusCode())
	contentType := ctx.Response().Header.Peek("Content-Type")
	assert.Equal(t, "application/json", string(contentType))

	// decode the response body and assert that it is equal to the error message
	var result map[string]string
	err = json.Unmarshal(ctx.Response().Body(), &result)
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"error": "Bad Request"}, result)
}
