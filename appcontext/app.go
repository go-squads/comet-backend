package appcontext

import(
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/go-squads/comet-backend/config"
)

type appContext struct {
	db *sql.DB
}

var context *appContext

func initializeDB() *sql.DB{
	db, err := sql.Open("postgres", config.ConnectionString())

	if err != nil{
		log.Fatalf(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf(err)
	}

	return db
}

func Initiate(){
	db := initializeDB()

	context = &appContext{
		db: db,
	}
}

func GetDB() *sql.DB{
	return context.db
}
