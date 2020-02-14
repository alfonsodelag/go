package server

import (
	"go/pkg/headers"
	"go/pkg/html"
	"go/pkg/json"
	textreq "go/pkg/text"
	"go/pkg/xml"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Routing comment: function for routing
func Routing() {
	router := mux.NewRouter()
	router.HandleFunc("/html", html.GetHTML).Methods("GET")
	router.HandleFunc("/json", json.GetJSON).Methods("GET")
	router.HandleFunc("/xml", xml.GetXML).Methods("GET")
	router.HandleFunc("/header", headers.GetHeader).Methods("GET")
	router.HandleFunc("/txt", textreq.GetTXT).Methods("GET")
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/img/", http.StripPrefix("/img/", fs))

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
