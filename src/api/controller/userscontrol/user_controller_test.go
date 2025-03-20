package userscontrol

import (
	"fmt"
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

	tests := []struct {
		name               string
		email              string
		statusCodeExpected int
	}{
		{"John Doe", "john0@doe.com", 201},
		{"John", "john1@doe.com", 201},
		// TODO - Implement the logic that will assert the success json in the correct case (new type for each case??)
		// {"Jo", "john2@doe.com", 400},
		// {"John Doe", "john3", 400},
		// {"", "", 400},
	}

	for _, tc := range tests {
		bodyRequest := fmt.Sprintf(
			`{
				"name": "%s",
				"email": "%s",
				"password": "example"
			}`, tc.name, tc.email,
		)
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

		assert.Equal(t, tc.statusCodeExpected, resp.StatusCode)
		assert.JSONEq(
			t,
			string(bodyResponse),
			fmt.Sprintf(
				`{
					"message": "User created successfully",
					"status": "success",
					"data": {
						"name": "%s",
						"email": "%s",
						"role": "user"
					}
				}`, tc.name, tc.email,
			),
		)
	}

}
