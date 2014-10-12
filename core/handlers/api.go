package handlers

import (
	"github.com/diasf/repogo/core/env"
	"github.com/diasf/repogo/core/net/mux"
)

var apiRouter = mux.NewRouterPath("/api/")

func init() {
	env.RegisterHandler(env.Handler{
		URLPrefix: apiRouter.GetPath(),
		Handler:   apiRouter,
	})
}

func RegisterAPIComponent(prefix string) mux.SubRouter {
	return apiRouter.SubRouter(prefix)
}
