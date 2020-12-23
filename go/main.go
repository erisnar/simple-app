package main

import "net/http"

func main() {
	// setup webserver
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}
