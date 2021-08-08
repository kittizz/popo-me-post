package api

import (
	"github.com/gofiber/fiber/v2"
)

type API struct {
	Fiber *fiber.App
	Addr  string
}

func NewAPI() (*API, error) {
	app := fiber.New()

	return &API{
		Fiber: app,
	}, nil
}
func (app *API) InjectRoute() {

}
