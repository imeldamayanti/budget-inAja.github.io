package config

//from github
import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(){
	// parseTime error : unsupported scan
	db, err := sql.Open("mysql", "root:root@/products?parseTime=true")

	if err != nil {
		panic(err)
	}

	//db in nampung koneksi db kita, bisa diakses
	log.Println("Database Connected")
	DB = db
}