package handler

import (
	"testing"
	"net/http/httptest"
	"bytes"
	"log"
	"net/http"

	"github.com/stretchr/testify/assert"
)



func TestCreateNewApplicationShouldReturnOk(t *testing.T)  {
	application := []byte(`{"app_name": "GO-JEK"}`)


	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST","/application/create", bytes.NewBuffer(application))
	if err != nil {
		log.Println(err.Error())
	}

	InsertNewApplication(w,req)
	assert.Equal(t, "{\"status\":200,\"message\":\"Inserted New Application\"}\n", w.Body.String())
}

func TestCreateNewApplicationShouldReturnBadRequest(t *testing.T)  {
	application := []byte(`{"app_name": "GO-JEK"}`)


	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST","/application/create", bytes.NewBuffer(application))
	if err != nil {
		log.Println(err.Error())
	}

	InsertNewApplication(w,req)
	assert.Equal(t, "{\"status\":400,\"message\":\"Duplicate Application Name\"}\n", w.Body.String())
}


func TestCreateNewApplicationWithoutAuthentication(t *testing.T)  {
	application := []byte(`{"app_name": "GO-JEK"}`)


	w := httptest.NewRecorder()
	handler := http.HandlerFunc(InsertNewApplication)

	req, err := http.NewRequest("POST","/application/create", bytes.NewBuffer(application))
	if err != nil {
		log.Println(err.Error())
	}

	//header := req.Header.Get("Authorization")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "")
	handler.ServeHTTP(w,req)

	assert.Equal(t, "{\"status\":401,\"message\":\"User Unauthorized\"}\n", w.Body.String())
}
