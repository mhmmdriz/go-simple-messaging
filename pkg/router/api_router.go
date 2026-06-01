package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/kooroshh/fiber-boostrap/app/controllers"
)

type ApiRouter struct{}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	api.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.Post("/register", controllers.Register)
	user.Post("/login", controllers.Login)

}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
