package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var db *sql.DB
var e error

func main() {
	fmt.Printf("gateway started\n")
	fmt.Printf("gateway started\n")
	fmt.Printf("gateway started\n")
	r := mux.NewRouter()

	db, e = sql.Open(
		"mysql",
		os.ExpandEnv("root:${MYSQL_SERVER_PASSWORD}@mysql_server:3306/${MYSQL_DATABASE}"),
	)
	fmt.Print("error is", e)

	r.HandleFunc("/todos", getTodos).Methods("GET")

	fmt.Printf("gateway started")
	http.ListenAndServe(":8080", r)
}

type todo struct{}

func getTodos(w http.ResponseWriter, r *http.Request) {
	t := new(todo)
	s, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, string(s))
}
