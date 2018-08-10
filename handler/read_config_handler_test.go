package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestReadConfig(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/configuration?app=comet_test&namespace=dev", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	ReadConfigurationHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"namespaceId\":0,\"version\":0,\"configurations\":null}\n", w.Body.String())
}
