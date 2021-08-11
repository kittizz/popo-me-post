//go:generate goversioninfo -icon=icons/favicon.ico -manifest=goversioninfo.exe.manifest

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
	"github.com/kittizz/popo-me-post/internal/pkg/api"
	"github.com/kittizz/popo-me-post/internal/pkg/config"
	"github.com/kittizz/popo-me-post/internal/pkg/ui"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

var quit = make(chan os.Signal, 1)

//go:embed init
var initUI embed.FS

//go:embed ui/dist/*
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

	c.Invoke(func(api *api.API, cfg *config.Config) {
		port := "0"
		if cfg.GetString("ENV") == config.DEVELOPMENT {
			port = "112"
		}
		ln, err := net.Listen("tcp", "127.0.0.1:"+port)
		if err != nil {
			logx.Fatal(err)
		}

		api.Addr = ln.Addr().String()

		if cfg.GetString("ENV") == config.PRODUCTION {
			api.Fiber.Use(filesystem.New(filesystem.Config{
				Root:       http.FS(&EmbedFS{dist}),
				PathPrefix: "vue/dist",
				Browse:     true,
			}))
		}

		go api.Fiber.Listener(ln)
	})

	c.Invoke(func(ui *ui.UI, api *api.API) *ui.UI {
		err := ui.Start(quit)
		if err != nil {
			logx.Fatal(err)
		}

		if viper.GetString("ENV") != config.DEVELOPMENT {
			ui.Host = api.Addr
		} else {
			ui.Host = "localhost:8080"
		}

		err = ui.Load(fmt.Sprintf("http://%s", ui.Host))
		if err != nil {
			logx.Fatal(err)
		}
		ui.Bind("ONxResize", ui.OnResize)
		ui.Bind("ONxLoadTab", ui.OnLoadTab)
		ui.LoadInitUI(&initUI)
		return ui
	})

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit

	c.Invoke(func(ui *ui.UI, api *api.API, c *config.Config) {
		logx.Println("saving config...")

		err := c.Save()
		if err != nil {
			logx.Println(err)
		}

		logx.Println("stopping UI...")

		err = ui.Close()
		if err != nil {
			logx.Println(err)
		}
		err = os.RemoveAll(ui.Dir)
		if err != nil {
			logx.Println(err)
		}

		logx.Println("stopping API...")
		api.Fiber.Server().CloseOnShutdown = true
		err = api.Fiber.Shutdown()
		if err != nil {
			logx.Println(err)
		}
	})
	logx.Println("exiting...")
}

type EmbedFS struct {
	f embed.FS
}

func (embed *EmbedFS) Open(name string) (fs.File, error) {
	if strings.HasSuffix(name, "/") {
		name += "index.html"
	}
	return embed.f.Open(name)
}
