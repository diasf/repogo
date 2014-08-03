package handlers

import (
	"fmt"
	"net/http"

	"github.com/diasf/repogo/core/net/mux"
)

var contentRouter mux.Router

func init() {
	contentRouter = RegisterAPIComponent("content").Handle("/", &ContentHandler{})
}

type ContentHandler struct {
}

func (h *ContentHandler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprint(rw, "Content ...")
}
