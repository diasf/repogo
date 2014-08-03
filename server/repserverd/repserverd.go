package main

import (
	"log"
	"net/http"

	"github.com/diasf/repogo/core/env"

	_ "github.com/diasf/repogo/client/webui"
	_ "github.com/diasf/repogo/core/handlers"
)

func main() {
	log.Println("Starting RePoGo")

	router := http.NewServeMux()
	for _, h := range env.GetHandlers() {
		router.Handle(h.URLPrefix, h.Handler)
	}

	http.ListenAndServe(":8000", router)
}
