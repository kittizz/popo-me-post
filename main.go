//go:generate goversioninfo -icon=icons/favicon.ico -manifest=goversioninfo.exe.manifest

package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/kittizz/popo-me-post/internal/pkg/api"
	"github.com/kittizz/popo-me-post/internal/pkg/config"
	"github.com/kittizz/popo-me-post/internal/pkg/ui"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

var quit = make(chan os.Signal, 1)

//go:embed init
var initUI embed.FS

//go:embed vue/dist/*
var dist embed.FS

var logx = log.New(os.Stderr, "[MAIN] ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		logx.Fatalln("app: cannot load location")
	}
	time.Local = loc

	c := dig.New()

	if err := c.Provide(config.NewConfig); err != nil {
		logx.Fatalln("app: cannot provide Config")
	}

	if err := c.Provide(api.NewAPI); err != nil {
		logx.Fatalln("app: cannot provide API")
	}

	if err := c.Provide(ui.NewUI); err != nil {
		logx.Fatalln("app: cannot provide UI")
	}

	c.Invoke(func(cfg *config.Config) {
		cfg.LoadConfig()
	})
	c.Invoke(func(api *api.API) *api.API {
		ln, err := net.Listen("tcp", "127.0.0.1:112")
		if err != nil {
			logx.Fatal(err)
		}

		api.Addr = ln.Addr().String()

		api.Fiber.Use(filesystem.New(filesystem.Config{
			Root:       http.FS(&EmbedFS{dist}),
			PathPrefix: "vue/dist",
			Browse:     true,
		}))
		go api.Fiber.Listener(ln)
		return api
	})

	c.Invoke(func(ui *ui.UI, api *api.API) {
		err := ui.Start(quit)
		if err != nil {
			logx.Fatal(err)
		}

		load := "http://localhost:8080"
		if viper.GetString("mode") != "dev" {
			load = fmt.Sprintf("http://%s", api.Addr)
		}
		err = ui.UI.Load(load)
		if err != nil {
			logx.Fatal(err)
		}

		ui.LoadInitUI(&initUI)

	})

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	c.Invoke(func(c *config.Config) {
		logx.Println("saving config...")

		err := c.Save()
		if err != nil {
			logx.Println(err)
		}
	})

	c.Invoke(func(ui *ui.UI) {
		logx.Println("stopping UI...")

		err := ui.UI.Close()
		if err != nil {
			logx.Println(err)
		}
	})

	c.Invoke(func(api *api.API) {
		logx.Println("stopping API...")
		api.Fiber.Server().CloseOnShutdown = true
		err := api.Fiber.Shutdown()
		if err != nil {
			logx.Println(err)
		}

	})

	logx.Println("exiting...")
}
