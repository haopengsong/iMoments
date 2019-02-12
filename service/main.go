package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Location Location `json:"location"`
}

const (
	DISTANCE = "200KM"
)

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")
	fmt.Println("get paras are: ", r.URL.Query())
	r.ParseForm()
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	num, _ := strconv.ParseFloat(r.URL.Query().Get("num"), 64)
	//radius := r.Form.Get("num")

	//range is optional
	ran := DISTANCE

	fmt.Println(num)
	if val := r.URL.Query().Get("num"); val != "" {
		ran = val + "km"
	}
	fmt.Println("range is ", ran)

	//return a fake post
	p := &Post{
		User:    "1111",
		Message: "working hard",
		Location: Location{
			Lat: lat,
			Lon: lon,
		},
	}

	jsonObj, err := json.Marshal(p)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonObj)
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
