package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contents, err := ioutil.ReadFile("app.go")
	if err != nil {
		fmt.Println(contents, err)
		return
	}

	for i:= 0; i < len(contents); i++ {
		print(string(contents[i]))
	}
}
