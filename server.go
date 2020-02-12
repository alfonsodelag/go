package server
import (
	"go_pill/pkg/headers-request"
	"go_pill/pkg/html-request"
	"go_pill/pkg/json-request"
	textreq "go_pill/pkg/text-request"
	"go_pill/pkg/xml-request"
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
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}