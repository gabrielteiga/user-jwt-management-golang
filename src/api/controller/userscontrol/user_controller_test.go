package userscontrol

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gabrielteiga/user-management-jwt/src/domain/services/userservice"
	"github.com/gabrielteiga/user-management-jwt/src/infrastructure/repositories"
	"github.com/gabrielteiga/user-management-jwt/src/infrastructure/repositories/userrepository"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

const CREATE_USER_ENDPOINT = "/api/v1/users"

func setupApp() *fiber.App {
	mysqlDB := repositories.MySQLDB{}
	err := mysqlDB.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	uc := NewUserController(
		userservice.NewUserService(
			userrepository.NewUserRepositorySQL(mysqlDB.DB),
		),
	)
	app.Post(CREATE_USER_ENDPOINT, uc.Create)

	return app
}

func TestUserCreate(t *testing.T) {
	app := setupApp()

	bodyRequest := `{
		"name": "John Doe",
		"email": "gabriel@lamberto2.com",
		"password": "password"
	}`
	req := httptest.NewRequest(http.MethodPost, CREATE_USER_ENDPOINT, strings.NewReader(bodyRequest))

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyResponse, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.JSONEq(
		t,
		string(bodyResponse),
		`{
    		"message": "User created successfully",
    		"status": "success",
    		"data": {
        		"name": "John Doe",
        		"email": "gabriel@lamberto2.com",
        		"role": "user"
    		}
		}`,
	)
}
