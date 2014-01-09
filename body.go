package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
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
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintln(w, string(body))
}
