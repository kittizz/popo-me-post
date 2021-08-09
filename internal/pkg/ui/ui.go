package ui

import (
	"embed"
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"

	"github.com/kittizz/popo-me-post/internal/pkg/config"
	"github.com/zserge/lorca"
)

type UI struct {
	UI  lorca.UI
	log *log.Logger

	config *config.Config
}

func NewUI(config *config.Config) *UI {
	return &UI{
		UI:     nil,
		log:    log.New(os.Stderr, "[UI] ", log.Ldate|log.Ltime|log.Lshortfile),
		config: config,
	}

}

func (u *UI) Start(c chan<- os.Signal) error {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	args = append(args, "--content-shell-hide-toolbar")
	args = append(args, "--disable-infobars ")
	args = append(args, "--disable-session-crashed-bubble")

	ui, err := lorca.New("", "", u.config.GetInt("windows-x"), u.config.GetInt("windows-y"), args...)
	if err != nil {
		return err
	}
	u.UI = ui

	go func() {
		<-ui.Done()
		c <- syscall.SIGINT
	}()

	return nil
}

func (u *UI) LoadInitUI(fs *embed.FS) {

	dirs, err := fs.ReadDir("init/javascript")
	if err != nil {
		u.log.Fatal(err)
	}

	for _, v := range dirs {
		if v.IsDir() {
			continue
		}

		u.log.Println(fmt.Sprintf("Load Init javascript %s", v.Name()))
		scr, err := fs.ReadFile("init/javascript/" + v.Name())
		if err != nil {
			u.log.Println("fail to load " + v.Name())
			continue
		}

		u.UI.Eval(string(scr))
	}
}
