package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Top)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.FormValue("uuid"))
}
