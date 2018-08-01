package handler

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {

	login := []byte(`{"Username": "comet", "Password": "comet_test"}`)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(login))

	if err != nil {
		log.Fatalf(err.Error())
	}

	LoginHandler(w, r)
	assert.Equal(t, "{\"status\":200,\"message\":\"log_in\",\"token\":\"rfBd67ti3SMtYvSgD6xAV1YU00zampta8Z8S686KLkIZ0PYkL28LTlsVqMNTZyLK\"}\n", w.Body.String())
}
