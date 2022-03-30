package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mokletdev/golang-fiber-codebase/bin"
	"github.com/mokletdev/golang-fiber-codebase/bin/config"
	"github.com/mokletdev/golang-fiber-codebase/bin/db"
	"github.com/mokletdev/golang-fiber-codebase/utils/res"
)

func main() {
	app := fiber.New()
	bin.Route(app)
	db.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(res.Response{Success: true, Data: "OK", Message: "Server is working properly", Code: http.StatusOK})
	})

	app.Listen(config.GlobalEnv.PORT)
}
