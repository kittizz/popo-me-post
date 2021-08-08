package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kittizz/popo-me-post/internal/pkg/api"
	"github.com/kittizz/popo-me-post/internal/pkg/ui"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

var quit = make(chan os.Signal, 1)

func main() {
	godotenv.Load(".env")
	viper.AutomaticEnv()
	viper.SetDefault("windows-x", 1433)
	viper.SetDefault("windows-y", 861)

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatalln("app: cannot load location")
	}
	time.Local = loc

	c := dig.New()

	if err := c.Provide(api.NewAPI); err != nil {
		log.Fatalln("app: cannot provide API")
	}

	if err := c.Provide(ui.NewUI); err != nil {
		log.Fatalln("app: cannot provide UI")
	}

	c.Invoke(func(api *api.API) *api.API {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}

		api.Addr = ln.Addr().String()

		go api.Fiber.Listener(ln)
		return api
	})

	c.Invoke(func(ui *ui.UI, api *api.API) {
		err := ui.Start(quit)
		if err != nil {
			log.Fatal(err)
		}
		err = ui.UI.Load(fmt.Sprintf("http://%s", api.Addr))
		if err != nil {
			log.Fatal(err)
		}

	})

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	c.Invoke(func(api *api.API) {
		log.Println("stopping API...")

		err := api.Fiber.Shutdown()
		if err != nil {
			log.Println(err)
		}
	})

	c.Invoke(func(ui *ui.UI) {
		log.Println("stopping UI...")

		err := ui.UI.Close()
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("exiting...")
}
