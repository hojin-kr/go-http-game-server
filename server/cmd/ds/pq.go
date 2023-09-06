package ds

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	database = "server"
)

func GetClient() *sql.DB {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + database + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
