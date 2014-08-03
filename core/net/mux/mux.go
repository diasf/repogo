/*
	Url muxer can register handler for hierarchic urls.
*/

package mux

import (
	"log"
	"net/http"
	"strings"
)

type Router interface {
	http.Handler
	Handle(subPath string, handler http.Handler) Router
	SubRouter(subPath string) Router
	GetPath() string
}

func NewRouter() Router {
	return NewRouterPath("/")
}

func NewRouterPath(path string) Router {
	subR := &router{
		pathPrefix: path,
		serveMux:   http.NewServeMux(),
		subRouter:  []*router{},
	}
	return subR
}

type router struct {
	pathPrefix string
	serveMux   *http.ServeMux
	subRouter  []*router
}

func (r *router) Handle(subPath string, handler http.Handler) Router {
	path := buildPath(r.pathPrefix, subPath)
	r.serveMux.Handle(path, handler)
	return r
}

func (r *router) SubRouter(subPath string) Router {
	path := buildPath(r.pathPrefix, subPath)
	subR := &router{
		pathPrefix: path,
		serveMux:   http.NewServeMux(),
		subRouter:  []*router{},
	}
	log.Printf("New subrouter for %s (%p) with prefix %s", r.pathPrefix, r, path)
	r.subRouter = append(r.subRouter, subR)
	return subR
}

func (r *router) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	uri := rq.URL.RequestURI()
	log.Printf("router %s (%p) called for %s and has %v subrouters", r.pathPrefix, r, uri, len(r.subRouter))
	if srvRouter := r.resolveRouter(uri); srvRouter != nil {
		srvRouter.ServeHTTP(w, rq)
	} else {
		log.Println("No router found for: ", uri)
	}
}

func (r *router) resolveRouter(uri string) *http.ServeMux {
	if strings.HasPrefix(uri, r.pathPrefix) {
		log.Printf("uri %s matched prefix %s ", uri, r.pathPrefix)
		for _, subR := range r.subRouter {
			if srvRouter := subR.resolveRouter(uri); srvRouter != nil {
				return srvRouter
			}
		}
		return r.serveMux
	}
	return nil
}

func (r *router) GetPath() string {
	return r.pathPrefix
}

func buildPath(root, child string) string {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	child = strings.TrimPrefix(child, "/")
	return root + child
}
