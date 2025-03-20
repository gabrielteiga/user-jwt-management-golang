package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

const HEALTH_ENDPOINT = "/api/v1/health"

func setupApp() *fiber.App {
	app := fiber.New()

	app.Get(HEALTH_ENDPOINT, Health)

	return app
}

func TestHealthIntegration(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, HEALTH_ENDPOINT, nil)

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, "The app is healthy!", string(body))
}
