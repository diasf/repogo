package handlers

import (
	"fmt"
	"net/http"

	"github.com/diasf/repogo/core/net/mux"
)

func init() {
	configureRoutes(RegisterAPIComponent("content"))
}

func configureRoutes(router mux.SubRouter) {
	ch := &contentHandler{}
	router.Method("GET").Handle("{contentId}", ch.findById)
}

type contentHandler struct {
}

func (h *ContentHandler) findById(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprint(rw, "Content ...")
}
