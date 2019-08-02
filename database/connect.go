package database

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conn() (db *sql.DB) {
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:8889)/user?parseTime=true")
	if err != nil {
	  fmt.Println("failed to connect database")
	}
	
	return db
}