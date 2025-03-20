package cmd

import (
	"log"

	"github.com/gabrielteiga/user-management-jwt/src/api/controller/userscontrol"
	"github.com/gabrielteiga/user-management-jwt/src/api/routes"
	"github.com/gabrielteiga/user-management-jwt/src/domain/services/userservice"
	"github.com/gabrielteiga/user-management-jwt/src/infrastructure/repositories"
	"github.com/gabrielteiga/user-management-jwt/src/infrastructure/repositories/userrepository"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	mysqlDB := repositories.MySQLDB{}
	mysqlDB.Open()
	defer mysqlDB.Close()

	jr := routes.NewJobRouter(
		app,
		userscontrol.NewUserController(
			userservice.NewUserService(
				userrepository.NewUserRepositorySQL(mysqlDB.DB),
			),
		),
	)
	jr.CreateRoutes()

	log.Fatal(app.Listen(":3001"))
}
