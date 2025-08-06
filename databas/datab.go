package databas

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	connstr := "user=postgres dbname=postgres password=pstgrs host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		fmt.Println("Hren tam")
	}
	DB = db
	return db
}
