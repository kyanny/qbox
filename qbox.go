package main

import (
	"fmt"
//	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"labix.org/v2/mgo"
//	"labix.org/v2/mgo/bson"
	"github.com/nu7hatch/gouuid"
)

type Content struct {
	Uuid string
//	Name string
}

func main() {
	http.HandleFunc("/", Top)
	http.HandleFunc("/new", Register)
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/contents", Serve)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Top(w http.ResponseWriter, r *http.Request) {
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		panic("Method Not Allowed")
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("contents")

	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	content := Content{uuid.String()}
	err = c.Insert(&content)
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(json))
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		panic("Method Not Allowed")
	}

	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	s3client := s3.New(auth, aws.USEast)
	bucket := s3client.Bucket("kyanny")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	uuid := r.FormValue("uuid")

	err = bucket.Put(uuid, data, "application/octet-stream", s3.BucketOwnerFull)
	if err != nil {
		panic(err)
	}
}

func Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		panic("Method Not Allowed")
	}

	uuid := r.FormValue("uuid")
	
	http.Redirect(w, r, "https://s3.amazonaws.com/kyanny/" + uuid, 302)
}
