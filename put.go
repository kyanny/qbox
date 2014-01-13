package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", Top)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Top(w http.ResponseWriter, r *http.Request) {
	fmt.Println("> " + time.Now().Format(time.RFC3339Nano))
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("  " + time.Now().Format(time.RFC3339Nano))
	fmt.Fprintf(w, "ok")
}
