package webui

import (
	"net/http"

	"github.com/diasf/repogo/core/env"
)

var (
	wEBUI_ROOT = "/home/fdias/go/src/github.com/diasf/repogo/client/webui/app"
	HandlerID  = "webui"
)

func init() {
	env.RegisterHandler(env.Handler{
		URLPrefix: "/",
		Handler:   http.FileServer(http.Dir(wEBUI_ROOT)),
	})
}
