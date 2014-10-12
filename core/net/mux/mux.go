/*
	Url muxer can register handler for hierarchic urls.
*/

package mux

import (
	"net/http"
	"regexp"
	"strings"
)

type Router interface {
	http.Handler
	Handler(handler http.Handler) Router
	Method(method string) Router
	Path(path string) Router
	SubRouter() Router
}

func NewRouter() Router {
	subR := &router{
		matchers:   []Matcher{},
		subRouters: []*router{},
	}
	return subR
}

type router struct {
	parent      *router
	handler     http.Handler
	pathMatcher *pathMatcher
	matchers    []Matcher
	subRouters  []*router
}

type Matcher interface {
	Match(*http.Request) MatchResult
}

type MatchResult interface {
	IsPartialMatch() bool
	IsFullMatch() bool
}

func (r *router) Method(method string) Router {
	// add method matcher
	return r
}

func (r *router) Path(path string) Router {
	var pathPattern = path
	if strings.Contains(path, "{") && strings.Contains(path, "}") {
		pathPattern = "[^/]+"
	}
	absolutePath := buildPath(r.buildRouterPath(), pathPattern)
	pattern := "^" + absolutePath + "/?"
	r.pathMatcher = &pathMatcher{
		path:        path,
		pathPattern: pathPattern,
		pattern:     pattern,
		regex:       regexp.MustCompile(pattern),
	}
	return r
}

func (r *router) buildRouterPath() string {
	parentPath := ""
	if r.parent != nil {
		parentPath = r.parent.buildRouterPath()
	}
	if r.pathMatcher != nil {
		parentPath = buildPath(parentPath, r.pathMatcher.pathPattern)
	}
	return parentPath
}

func (r *router) Handler(handler http.Handler) Router {
	r.handler = handler
	return r
}

func (r *router) SubRouter() Router {
	subR := NewRouter()
	subR.(*router).parent = r
	r.subRouters = append(r.subRouters, subR.(*router))
	return subR
}

func (r *router) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	var handler http.Handler
	if matchedRouter := r.matchRouter(rq); matchedRouter != nil {
		handler = matchedRouter.handler
	}

	if handler == nil {
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, rq)
}

func (r *router) matchRouter(rq *http.Request) *router {
	fullMatch := true
	if r.pathMatcher != nil {
		if mr := r.pathMatcher.Match(rq); !mr.IsPartialMatch() {
			return nil
		} else if !mr.IsFullMatch() {
			fullMatch = false
		}
	}
	for _, m := range r.matchers {
		if mr := m.Match(rq); !mr.IsPartialMatch() {
			return nil
		} else if !mr.IsFullMatch() {
			fullMatch = false
		}
	}
	for _, s := range r.subRouters {
		if matchedSub := s.matchRouter(rq); matchedSub != nil {
			return matchedSub
		}
	}

	if fullMatch {
		return r
	}

	return nil
}

func buildPath(root, child string) string {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	child = strings.TrimPrefix(child, "/")
	return root + child
}

type pathMatcher struct {
	path        string
	pathPattern string
	pattern     string
	regex       *regexp.Regexp
}

func (p *pathMatcher) Match(rq *http.Request) MatchResult {
	rs := p.regex.Split(rq.URL.Path, 2)
	partialMatch := len(rs) == 2 && len(rs[0]) == 0
	fullMatch := partialMatch && len(rs[1]) == 0
	return DefaulMatchResult{
		partialMatch: partialMatch,
		fullMatch:    fullMatch,
	}
}

type DefaulMatchResult struct {
	partialMatch bool
	fullMatch    bool
}

func (d DefaulMatchResult) IsPartialMatch() bool {
	return d.partialMatch
}

func (d DefaulMatchResult) IsFullMatch() bool {
	return d.fullMatch
}
