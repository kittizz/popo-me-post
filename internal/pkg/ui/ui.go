package ui

import (
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"syscall"

	"github.com/kittizz/lorca"
	"github.com/kittizz/popo-me-post/internal/pkg/config"
)

type UI struct {
	lorca.UI
	Dir  string
	Host string

	log    *log.Logger
	c      chan<- os.Signal
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
	args = append(args, "--kiosk-print")
	args = append(args, "--overscroll-history-navigation=0")
	args = append(args, "--disable-pinch")
	dir, err := ioutil.TempDir("", "popomepost")
	if err != nil {
		log.Fatal(err)
	}
	u.Dir = dir

	ui, err := lorca.New("", u.Dir, u.config.GetInt("windows-x"), u.config.GetInt("windows-y"), args...)
	if err != nil {
		return err
	}
	u.UI = ui
	u.c = c
	go func() {
		<-u.Done()
		u.Exit()
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

		u.AddScriptToEvaluateOnNewDocument(string(scr))
	}
}
func (u *UI) Exit() {
	u.c <- syscall.SIGINT
}
