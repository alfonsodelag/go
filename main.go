package main

import (
	"go_pill/pkg/headers"
	"go_pill/pkg/html"
	"go_pill/pkg/json"
	"go_pill/pkg/text"
	"go_pill/pkg/xml"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Routing comment: function for routing
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/html", html.Foo).Methods("GET")
	router.HandleFunc("/json", json.Foo).Methods("GET")
	router.HandleFunc("/xml", xml.Foo).Methods("GET")
	router.HandleFunc("/header", headers.ServerHTTP).Methods("GET")
	router.HandleFunc("/txt", text.Foo).Methods("GET")
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/img/", http.StripPrefix("/img/", fs))

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
