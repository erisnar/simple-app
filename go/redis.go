package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var client = rClient()
var ctx = context.Background()

// connect to redis database
func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		// redis-master endpoint is created
		Addr: "redis-master:6379",
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
