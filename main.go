package main

import (
	"database/sql"
	"flag"
	"github.com/rs/cors"
	"github.com/woodhds/vk.service/database"
	"github.com/woodhds/vk.service/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var token string
var version string
var count int

func main() {

	flag.StringVar(&token, "token", "", "access token required")
	flag.StringVar(&version, "version", "", "version required")
	flag.IntVar(&count, "count", 10, "used count")
	flag.Parse()

	if token == "" {
		panic("access token required")
	}

	log.Printf("Used token: %s", token)
	log.Printf("Used version: %s", version)
	log.Printf("Used count: %d", count)

	conn, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	database.Migrate(conn)

	r := mux.NewRouter()

	r.Path("/messages").Handler(handlers.MessagesHandler(conn)).Methods(http.MethodGet)

	r.Path("/grab").Handler(handlers.ParserHandler(conn, token, version, count)).Methods(http.MethodGet)

	r.Path("/users").Handler(handlers.UsersHandler(conn)).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	r.Path("/repost").Handler(handlers.RepostHandler(token, version)).Methods(http.MethodPost, http.MethodOptions)

	r.Path("/users/search").Handler(handlers.UsersSearchHandler(token, version)).Methods(http.MethodGet, http.MethodOptions)

	http.ListenAndServe(":4222", cors.Default().Handler(r))
}
