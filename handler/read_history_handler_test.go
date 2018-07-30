package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadHistoryConfig(t *testing.T){

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET","/history/comet_test/dev",nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	ReadHistoryConfiguration(w,r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

}
