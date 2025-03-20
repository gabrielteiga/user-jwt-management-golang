package routes

import (
	"github.com/gabrielteiga/user-management-jwt/src/api/controller"
	"github.com/gabrielteiga/user-management-jwt/src/api/controller/userscontrol"
	"github.com/gofiber/fiber/v2"
)

type JobRouter struct {
	Router         *fiber.App
	UserController *userscontrol.UserController
}

func NewJobRouter(router *fiber.App, uc *userscontrol.UserController) *JobRouter {
	return &JobRouter{
		Router:         router,
		UserController: uc,
	}
}

func (jr *JobRouter) CreateRoutes() {
	api := jr.Router.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/health", controller.Health)
	v1.Post("/users", jr.UserController.Create)
}
