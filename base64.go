package main

import (
	"fmt"
	"time"
	"encoding/base64"
)

func main() {
	t := time.Now().Format(time.RFC3339Nano)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(t)))
}
