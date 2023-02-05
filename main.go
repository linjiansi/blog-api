package main

import (
	"database/sql"
	"fmt"
	"github.com/linjiansi/blog-api/router"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbUser     = os.Getenv("USERNAME")
	dbPassword = os.Getenv("USERPASS")
	dbDatabase = os.Getenv("DATABASE")
	dbConn     = fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser,
		dbPassword,
		dbDatabase,
	)
)

func main() {

	db, err := connectDB()
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := router.NewRouter(&db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
