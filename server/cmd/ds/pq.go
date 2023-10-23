package ds

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	dbname   = os.Getenv("DB_NAME")
)

func GetClient() *sql.DB {
	log.Println("DB_USER: " + user)
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
