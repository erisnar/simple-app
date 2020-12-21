package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v8"
)

type rObject struct {
	Key   string
	Value string
}

func rObjectCreate(res http.ResponseWriter, req *http.Request) {
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

var client = rClient()
var ctx = context.Background()

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func handler(res http.ResponseWriter, req *http.Request) {

	// Example of parsing GET or POST Query Params.
	req.ParseForm()

	// Example of handling POST request.
	switch req.Method {
	case "POST":
		rObjectCreate(res, req)
	// Example of handling GET request.
	case "GET":
		get(res, req)
	default:
		bad(res, req)
	}
}

func get(res http.ResponseWriter, req *http.Request) {
	// Example of fetching specific Query Param.
	fmt.Fprintf(res, "Getting value of key: %s\n", req.URL.Path)
	var trimPath = strings.Trim(req.URL.Path, "/")
	var value = getValue(trimPath)

	fmt.Fprintf(res, "Value: %s\n", value)
}

func bad(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "404", req.Method, req.URL.Path)
}

func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}

// add key value
func setValue(key string, value string) error {

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
	} else {
		fmt.Println("Created: ", key, value)
	}

	return nil
}

// get value of key
func getValue(key string) string {

	fmt.Println("Getting: ", key)

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("404 Not found")

		val = "404 Not found"
	} else {
		fmt.Println("Returned: ", val)
	}

	return val
}
