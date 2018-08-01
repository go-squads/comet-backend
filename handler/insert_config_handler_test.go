package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-squads/comet-backend/appcontext"
	"github.com/go-squads/comet-backend/domain"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	appcontext.Initiate()
	testResult := m.Run()
	os.Exit(testResult)
}

func TestInsertConfigHandlerReturns400WhenBodyFormatIsIncorrect(t *testing.T) {
	var response domain.Response
	reqBody := []byte(`{"Username": "comet", "Password": "comett"}`)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/configuration", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatalf(err.Error())
	}

	InsertConfigurationHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, http.StatusBadRequest, response.Status)
}

func TestInsertConfigHandlerReturns400WhenAppIsNotCreated(t *testing.T) {
	var response domain.Response
	reqBody := []byte(`{"appName": "comet_unavailable", "namespace": "new", "data": [{"key": "DBNAME", "value": "comet"}]}`)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/configuration", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatalf(err.Error())
	}

	InsertConfigurationHandler(w, r)
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, http.StatusBadRequest, response.Status)
}
