//go:generate goversioninfo -icon=icons/ms-icon-310x310.ico -manifest=goversioninfo.exe.manifest

package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
	"github.com/kittizz/popo-me-post/internal/pkg/api"
	"github.com/kittizz/popo-me-post/internal/pkg/ui"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

var quit = make(chan os.Signal, 1)

//go:embed init
var initUI embed.FS

//go:embed vue/dist/*
var dist embed.FS

type EmbedFS struct {
	f embed.FS
}

// append index.html to any 'directories' that are opened
func (embed *EmbedFS) Open(name string) (fs.File, error) {
	if strings.HasSuffix(name, "/") {
		name += "index.html"
	}
	return embed.f.Open(name)
}

var logx = log.New(os.Stderr, "[MAIN] ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	godotenv.Load(".env")
	viper.AutomaticEnv()
	viper.SetDefault("windows-x", 1433)
	viper.SetDefault("windows-y", 861)

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		logx.Fatalln("app: cannot load location")
	}
	time.Local = loc

	c := dig.New()

	if err := c.Provide(api.NewAPI); err != nil {
		logx.Fatalln("app: cannot provide API")
	}

	if err := c.Provide(ui.NewUI); err != nil {
		logx.Fatalln("app: cannot provide UI")
	}

	c.Invoke(func() {

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
		// err := api.Fiber.Shutdown()
		// if err != nil {
		// 	logx.Println(err)
		// }

	})
	logx.Println("exiting...")
}
