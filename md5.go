package main

import (
	"fmt"
	"time"
	"crypto/md5"
)

func main() {
	t := time.Now().Format(time.RFC3339Nano)
	h := md5.New()
	d := h.Sum([]byte(t))
	fmt.Printf("%x\n", d)
}
