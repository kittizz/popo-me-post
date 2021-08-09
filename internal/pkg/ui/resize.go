package ui

func (u *UI) OnResize(x int, y int) {

	u.config.Set("windows-x", x)
	u.config.Set("windows-y", y)
	u.log.Printf("resize x:%v,y:%v", x, y)
}
