package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	User string `json:"user"`
	Message string `json:"message"`
	Location Location `json:"location"`
}

func main() {
	fmt.Println("started-service")
    http.HandleFunc("/post", handlerPost)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
		return
	}
	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}