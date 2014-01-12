package main

import (
	"encoding/json"
	"os"
)

type Message struct {
	Id int64
	Data string
}

func main() {
	msg := Message{1, "JSONTest"}
	json, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(json)
}
