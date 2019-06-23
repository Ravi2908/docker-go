package main

import (
	"net/http"
	"testing"
)

func TestNewServer(t *testing.T) {
	mux := http.NewServeMux()
	if err := NewServer(mux); err != nil {
		t.Errorf("Unable to start the server")
	}
}
