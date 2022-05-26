package main

import (
	"go-bookstore-api/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "fmt"
// "log"
// "net/http"
//"github.com/gorilla/mux"
//"github.com/jinzhu/gorm/dialects/mysql"
//"go-bookstore-api/pkg/routes"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
