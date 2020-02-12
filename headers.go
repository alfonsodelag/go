package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("alfonsokey", "this is the key from alfonso")
	fmt.Fprintln(res, "Any code you want in this func")
}

func main() {
	var h hotdog
	http.ListenAndServe(":3000", h)
}
