package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-squads/comet-backend/appcontext"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	appcontext.Initiate()
	testResult := m.Run()
	os.Exit(testResult)
}

func TestReadConfig(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/configuration?app=comet_test&namespace=dev", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	ReadConfigurationHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "[{\"namespaceId\":1,\"version\":1,\"key\":\"PGUSERNAME\",\"value\":\"postgres\"}]\n", w.Body.String())
}
