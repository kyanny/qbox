package main

import (
	"fmt"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

func main() {
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	s3client := s3.New(auth, aws.USEast)
	bucket := s3client.Bucket("kyanny")
	fmt.Println(bucket)

	data := []byte("hello, world")
	err = bucket.Put("hoge.txt", data, "text/plain", s3.BucketOwnerFull)
	if err != nil {
		panic(err.Error())
	}

	content, err := bucket.Get("hoge.txt")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(content))
}
