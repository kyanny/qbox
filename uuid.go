package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

func main() {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	fmt.Println(uuid)
	uuid_str := uuid.String()
	println(uuid_str)
}
