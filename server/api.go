package server

import (
	"log"
	"net/http"
)

func StartAPIServer() {
	muxRouter := Router()
	log.Fatal(http.ListenAndServe(":8000", muxRouter))
}
