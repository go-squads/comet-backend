package main

import (
	"github.com/go-squads/comet-backend/appcontext"
	"github.com/go-squads/comet-backend/server"
)

func main() {
	appcontext.Initiate()
	server.StartAPIServer()
}
