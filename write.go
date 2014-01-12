package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contents, err := ioutil.ReadFile("/Users/kyanny/tmp/zero.dat")
	if err != nil {
		fmt.Println(contents, err)
		return
	}

	println("writing...")
	ioutil.WriteFile("zero.dat", contents, 0644)
	println("done.")
}
