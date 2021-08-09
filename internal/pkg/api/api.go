package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	Fiber *fiber.App
	Addr  string
	log   *log.Logger
}

func NewAPI() (*API, error) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ServerHeader:          "Fiber",
		AppName:               "POPO ME POST",
	})
	return &API{
		Fiber: app,
		log:   log.New(os.Stderr, "[API] ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}
func (app *API) InjectRoute() {

}
