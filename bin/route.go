package bin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mokletdev/golang-fiber-codebase/bin/modules/User"
)

func Route(app *fiber.App) {
	app.Get("/api/users", User.GetUsers)
	app.Post("api/users", User.CreateUser)
}
