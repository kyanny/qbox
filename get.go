package main

import (
	"fmt"
	"net/http"
	"bytes"
	"os"
)

func main() {
	http.HandleFunc("/", Top)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Top(w http.ResponseWriter, r *http.Request) {
	method := bytes.ToUpper([]byte(r.Method))
	if string(method) != "GET" {
		http.Error(w, string(method) + " is not allowed", 403)
	} else {
		fmt.Fprintln(w, r.Method + " is ok")
	}
}
