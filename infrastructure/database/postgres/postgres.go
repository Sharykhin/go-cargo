package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // import postgers dependencies
)

// NewConnection establishes a new connection with Postgres
func NewConnection(user, password, host, dbname string, port int) *sql.DB {
	source := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", source)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	return db
}
