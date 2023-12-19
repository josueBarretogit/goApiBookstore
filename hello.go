package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		fmt.Printf("hello request from go %s \n", req.URL)
	})
	http.ListenAndServe(":80", nil)
}
