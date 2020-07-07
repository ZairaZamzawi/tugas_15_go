package main

import (
	"fmt"
	"net/http"
)

func index(web http.ResponseWriter, r *http.Request) {
	for i := 1; i < 101; i++ {
		fmt.Fprintln(web, "Angka", i)
	}
}

func main() {
	http.HandleFunc("/index", index)

	// start web server
	http.ListenAndServe(":90", nil)
}
