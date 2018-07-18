package appcontext

import (
	"database/sql"
	"log"

	"github.com/go-squads/comet-backend/config"
	_ "github.com/lib/pq"
)

type appContext struct {
	db *sql.DB
}

var context *appContext

func initializeDB() *sql.DB {
	db, err := sql.Open("postgres", config.ConnectionString())

	if err != nil {
		log.Fatalf("%s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("%s", err)
	}

	return db
}

func Initiate() {
	db := initializeDB()

	context = &appContext{
		db: db,
	}
}

func GetDB() *sql.DB {
	return context.db
}
