package headers

import (
	"fmt"
	"net/http"
)

type hotdog int

//ServerHTTP function
func ServerHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("alfonsokey", "this is the key from alfonso")
	fmt.Fprintln(res, "Any code you want in this func")
}

func main() {
	http.HandleFunc("/", ServerHTTP)
	http.ListenAndServe(":3000", nil)
}
