package main

import (
	"database/sql"
	"fmt"
	"go_fullstack/handlers"
	"go_fullstack/store"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://gotest_owner:MdQoNmu8nS1A@ep-soft-tree-a16kjsj3.ap-southeast-1.aws.neon.tech/gotest?sslmode=require"
	db, err := store.NewMySQLStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}
	currStore := store.NewStore(db)
	store.InitDb(db)
	initStorage(db)

	router := mux.NewRouter()
	handler := handlers.New(currStore)

	router.HandleFunc("/", handler.HandleHome).Methods("GET")
	router.HandleFunc("/todo", handler.HandleTodo).Methods("GET")
    router.HandleFunc("/todo", handler.AddTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", handler.HandleSearchTodo).Methods("POST")
    router.HandleFunc("/todo/{id}", handler.HandleDeleteTodo).Methods("DELETE")
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", router)

}

func initStorage(db *sql.DB) {
	fmt.Println("initStorage")
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("Connected to Database!")
}
