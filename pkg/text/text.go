package text

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Foo)
	http.ListenAndServe(":3000", nil)
}

//Foo function
func Foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
