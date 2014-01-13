package main

import (
	"fmt"
	"time"
	"encoding/hex"
)

func main() {
	t := time.Now().Format(time.RFC3339Nano)
	fmt.Println(hex.EncodeToString([]byte(t)))
}
