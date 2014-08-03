package handlers

import (
	"fmt"
	"net/http"
)

func init() {
	RegisterAPIComponent("version").Handle("/", &versionHandler{})
}

type versionHandler struct {
}

func (h *versionHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprint(rw, "Api version 0.1")
}
