package config

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "myactionagiledb.ciowzhlnuioi.eu-west-3.rds.amazonaws.com"
	port     = 5432
	user     = "oconq"
	password = "tebfAz-qehvig-7roxga"
	dbname   = "dbmyfinance"
)

func DatabaseInit() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disaable", host, port, user, password, dbname)

	fmt.Printf("Connecting database...")

	//db, err = sql.Open("postgres", "jdbc:postgresql://myactionagiledb.ciowzhlnuioi.eu-west-3.rds.amazonaws.com:5432/dbmyfinance")
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("FAILED!")
		panic(err)
	}

	fmt.Println("OK")
}

func CheckAccountsTable() {
	_, err := db.Exec("SELECT * FROM Accounts")

	if err != nil {
		log.Fatal(err)
	}
}

func Db() *sql.DB {
	return db
}