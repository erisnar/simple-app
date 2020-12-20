package main

import (
	"fmt"
	"log"
    "context"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"
)

var ctx = context.Background()


func main() {

	set("foo", "bar")

	get("foo")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}
	
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
  
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting value of key: %s\n", r.URL.Path)

	var trimPath = strings.Trim(r.URL.Path, "/")
	var value = get(trimPath)

	fmt.Fprintf(w, "Value: %s\n", value)
}


// add key value
func set(key string, value string) {
	
	// connect to redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})

	err := rdb.Set(ctx, key, value, 0).Err()
    if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Created: ", key, value)
}

// get value of key
func get(key string) (string) {

	// connect to redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})

	fmt.Println("Getting: ", key)

    val, err := rdb.Get(ctx, key).Result()
    if err != nil {
		log.Fatal(err)
        panic(err)
	}
	
	fmt.Println("Returned: ", val)

	return val
}