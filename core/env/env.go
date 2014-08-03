package env

import (
	"net/http"
)

var envInstance = env{
	handlers: []Handler{},
}

type env struct {
	handlers []Handler
}

func RegisterHandler(handler Handler) {
	envInstance.handlers = append(envInstance.handlers, handler)
}

func GetHandlers() []Handler {
	return envInstance.handlers
}

type Handler struct {
	URLPrefix string
	Handler   http.Handler
}
