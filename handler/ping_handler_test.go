package handler

import (
	"net/http"
	"net/http/httptest"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandlert (t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	PingHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"success\": \"pong\"}", w.Body.String())
}
