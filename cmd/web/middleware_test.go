package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	handler := NoSurf(&myH)

	switch v := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	handler := SessionLoad(&myH)

	switch v := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}
