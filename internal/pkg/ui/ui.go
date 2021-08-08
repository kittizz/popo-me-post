package ui

import (
	"os"
	"runtime"
	"syscall"

	"github.com/spf13/viper"
	"github.com/zserge/lorca"
)

type UI struct {
	UI lorca.UI
}

func NewUI() *UI {

	return &UI{
		UI: nil,
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
	ui, err := lorca.New("", "", viper.GetInt("windows-x"), viper.GetInt("windows-y"), args...)
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
