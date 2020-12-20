package main

import (
	"fmt"
	"log"
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	
	// connect to redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})

	set("foo", "bar", rdb)

	get("foo", rdb)
}

// add key value
func set(key string, value string, rdb *redis.Client) {

	err := rdb.Set(ctx, key, value, 0).Err()
    if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Created: ", key, value)
}

// get value of key
func get(key string, rdb *redis.Client) (string) {

	fmt.Println("Getting: ", key)

    val, err := rdb.Get(ctx, key).Result()
    if err != nil {
		log.Fatal(err)
        panic(err)
	}
	
	fmt.Println("Returned: ", val)

	return val
}