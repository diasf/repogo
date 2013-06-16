package controllers

import (
	"github.com/robfig/revel"
	"runtime"
)

type App struct {
	*revel.Controller
}

func init() {
	revel.INFO.Println("Starting RepoGo...")
	revel.OnAppStart(func() {
		runtime.GOMAXPROCS(runtime.NumCPU())
	})
}

func (c App) Index() revel.Result {
	return c.Redirect("/ext/admin/index.html")
}
