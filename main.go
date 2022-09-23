package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ivanauliaa/learn-wire/user"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userHandler := user.Wire(db)
	http.Handle("/users", userHandler.FetchByUsername())

	log.Printf("listening on %s:%s\n", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
}
