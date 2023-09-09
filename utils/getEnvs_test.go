package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvs(t *testing.T) {
	// set the environment variables for testing
	os.Setenv("PORT", "8080")
	os.Setenv("DB_URL", "postgres://user:password@localhost:5432/mydb")

	// call the GetEnvs function
	result, err := GetEnvs()

	// assert that the result and error are as expected
	assert.Nil(t, err)
	assert.Equal(t, "8080", result.Port)
	assert.Equal(t, "postgres://user:password@localhost:5432/mydb", result.DB_URL)

	// unset the environment variables
	os.Unsetenv("PORT")
	os.Unsetenv("DB_URL")
}
