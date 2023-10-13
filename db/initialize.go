package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	goqu "github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
	db   *sql.DB
)

func InitializeDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}

func GetDB() *sql.DB {
	once.Do(func() {
		InitializeDatabase()
	})

	return db
}

func GetQuilderBuilder() goqu.DialectWrapper {
	dialect := goqu.Dialect("postgres")

	return dialect

}
