package main

import (
	"fmt"
	"time"
	"crypto/sha512"
)

func main() {
	t := time.Now().Format(time.RFC3339Nano)
	h := sha512.New()
	d := h.Sum512([]byte(t))
	fmt.Printf("%x\n", d)
}
