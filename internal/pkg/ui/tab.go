package ui

func (u *UI) OnLoadTab() {
	if u.Eval("location.host").String() != u.Host {
		u.Eval("window.close();")
		u.Exit()
	}
	//location.href = "https://google.com"
}
