package cachehttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCacheHandlerEmptyKey(t *testing.T) {
	handler := http.HandlerFunc(CacheHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := "Empty key not allowed\n"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestCacheHandlerNotAllowed(t *testing.T) {
	handler := http.HandlerFunc(CacheHandler)

	req, _ := http.NewRequest("PUT", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

	expected := "Method not allowed\n"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func TestCacheHandlerSuccess(t *testing.T) {

	handler := http.HandlerFunc(CacheHandler)

	reqPost, _ := http.NewRequest("POST", "/test", nil)
	recPost := httptest.NewRecorder()
	handler.ServeHTTP(recPost, reqPost)
	if status := recPost.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	reqGet, _ := http.NewRequest("GET", "/test", nil)
	recGet := httptest.NewRecorder()
	handler.ServeHTTP(recGet, reqGet)
	if status := recGet.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
