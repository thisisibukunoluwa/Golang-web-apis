package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/thisisibukunoluwa/Golang-web-apis/coffeester/db"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API is listening on port", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",port),
	}
	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := Config {
		Port: os.Getenv("PORT"),
	}
	dsn := os.Getenv("DSN")
	dbConn,err := db.ConnectPostgres(dsn)
	app := &Application{
		Config: cfg,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}