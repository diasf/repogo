package mux

import (
	"net/http"
	"testing"
)

func TestResolveRouter(t *testing.T) {
	testRouter := NewRouter().(*router)

	// build up the test data
	testRouter.Handle("/", mockHandler{})
	apiTestRouter := testRouter.SubRouter("/api").Handle("version", mockHandler{}).(*router)
	apiContentTestRouter := apiTestRouter.SubRouter("content").(*router)
	apiContentTestRouter.Handle("/download", mockHandler{})
	apiContentTestRouter.Handle("upload", mockHandler{})
	apiContentTestRouter.Handle("/", mockHandler{})

	// check whether the correct router is used for the specified url
	testUri(t, testRouter, "/", testRouter.serveMux)
	testUri(t, testRouter, "/api/version", apiTestRouter.serveMux)
	testUri(t, testRouter, "/api/content/download", apiContentTestRouter.serveMux)
	testUri(t, testRouter, "/api/content/upload", apiContentTestRouter.serveMux)
	testUri(t, testRouter, "/api/content/", apiContentTestRouter.serveMux)
}

func testUri(t *testing.T, testRouter *router, uri string, expectedRouter *http.ServeMux) {
	if testRouter.resolveRouter(uri) != expectedRouter {
		t.Errorf("Resolve %v failed", uri)
	}
}

type mockHandler struct {
}

func (h mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
