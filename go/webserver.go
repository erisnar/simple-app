package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// key value object stored in redis
type rObject struct {
	Key   string
	Value string
}

func handler(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	switch req.Method {
	case "POST":
		post(res, req)
	case "GET":
		get(res, req)
	default:
		bad(res, req)
	}
}

func post(res http.ResponseWriter, req *http.Request) {

	var r rObject

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}
	log.Println(r.Key)

	err = setValue(r.Key, r.Value)
	if err == nil {
		fmt.Fprintf(res, "Wrote %s:%s to database\n", r.Key, r.Value)
	} else {
		fmt.Fprintf(res, "Error")
	}
}

// get value of key from redis
func get(res http.ResponseWriter, req *http.Request) {

	// trim url
	var trimPath = strings.Trim(req.URL.Path, "/")
	fmt.Fprintf(res, "Getting value of key: %s\n", trimPath)

	// get value
	var value = getValue(trimPath)
	fmt.Fprintf(res, "Value: %s\n", value)
}

func bad(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "404\n")
}
