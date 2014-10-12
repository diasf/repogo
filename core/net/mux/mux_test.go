package mux

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testData struct {
	path          string
	requestPath   string
	ignoreRequest bool
	ignorePath    bool
	pushSub       bool
	popSub        bool
	attachHandler bool
	response      string
}

func TestSimpleRoutes(t *testing.T) {
	var testTable = []testData{
		{
			path:          "/api/version",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "version 1.1",
		},
		{
			path:          "/api/download",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "download it now",
		},
		{
			path:          "/api/upload",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "upload it now",
		},
		{
			path:       "/unknown/path",
			ignorePath: true,
			response:   "404 page not found",
		},
		{
			path:          "/other/thing/to/come",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "some other thing",
		},
	}

	executeTestTable(t, testTable)
}

func TestSubRoutes(t *testing.T) {
	var testTable = []testData{
		{
			path:     "/content",
			pushSub:  true,
			response: "404 page not found",
		},
		{
			path:          "download",
			requestPath:   "/content/download",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "download handler",
		},
		{
			path:          "upload",
			requestPath:   "/content/upload",
			pushSub:       true,
			attachHandler: true,
			response:      "upload handler",
		},
		{
			path:          "show",
			requestPath:   "/content/upload/show",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "content upload show",
		},
		{
			path:          "unknownsub",
			requestPath:   "/content/upload/wrongsub",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "404 page not found",
		},
		{
			ignorePath:    true,
			ignoreRequest: true,
			popSub:        true,
		},
		{
			path:          "show",
			requestPath:   "/content/show",
			pushSub:       true,
			attachHandler: true,
			response:      "content show",
		},
		{
			path:          "upload",
			requestPath:   "/content/show/upload",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "content show upload",
		},
		{
			ignorePath:    true,
			ignoreRequest: true,
			popSub:        true,
		},
		{
			ignorePath:    true,
			ignoreRequest: true,
			popSub:        true,
		},
		{
			path:          "/version",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "this is a version",
		},
	}

	executeTestTable(t, testTable)
}

func TestPattern(t *testing.T) {
	var testTable = []testData{
		{
			path:     "/content",
			pushSub:  true,
			response: "404 page not found",
		},
		{
			path:          "{category}",
			requestPath:   "/content/books",
			pushSub:       true,
			attachHandler: true,
			response:      "category",
		},
		{
			requestPath: "/content/cds",
			response:    "category",
		},
		{
			path:          "{id}",
			requestPath:   "/content/books/23",
			pushSub:       true,
			popSub:        true,
			attachHandler: true,
			response:      "book detail",
		},
		{
			requestPath: "/content/upload/show/bla",
			response:    "404 page not found",
		},
	}

	executeTestTable(t, testTable)
}

func executeTestTable(t *testing.T, table []testData) {
	testRouter := NewRouter()
	currentRouter := testRouter
	routerStack := []Router{}
	routerStack = append(routerStack, testRouter)
	for _, test := range table {
		if test.pushSub {
			routerStack = append(routerStack, currentRouter.SubRouter())
		}

		currentRouter = routerStack[len(routerStack)-1]

		if test.path != "" && !test.ignorePath {
			currentRouter.Path(test.path)
		}

		if test.attachHandler {
			currentRouter.Handler(buildHandler(test.response))
		}

		if test.popSub {
			routerStack = routerStack[:len(routerStack)-1]
			currentRouter = routerStack[len(routerStack)-1]
		}
	}

	for _, test := range table {
		if test.ignoreRequest {
			continue
		}
		requestPath := test.path
		if test.requestPath != "" {
			requestPath = test.requestPath
		}
		req, err := http.NewRequest("GET", requestPath, nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		testRouter.ServeHTTP(w, req)

		if rs := strings.TrimSpace(w.Body.String()); rs != test.response {
			t.Errorf("For test url '%s': Expected '%s' got '%s'", test.requestPath, test.response, rs)
		}
	}
}

func buildHandler(out string) mockHandler {
	return mockHandler{
		rs: out,
	}
}

type mockHandler struct {
	rs string
}

func (h mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.rs)
}
