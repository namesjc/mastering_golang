package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleGetVideos)
	http.HandleFunc("/update", HandleUpdateVideos)
	http.ListenAndServe(":8080", nil)
}

// func HandleGetVideos(w http.ResponseWriter, r *http.Request) {
// 	for header, value := range r.Header {
// 		fmt.Println("Key: %v \t Value: %v \n", header, value)
// 	}
// 	w.Header().Add("TestHeader", "TestValue")
// 	w.Write([]byte("Hello!"))
// }

func HandleGetVideos(w http.ResponseWriter, r *http.Request) {

	videos := getVideos()
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	w.Write(videoBytes)

}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var videos []video
		err= json.Unmarshal(body, &videos)
		if err != nil {
			w.WriteHeader(400)
			fmt.Println(w, "Bad request!")
		}

		saveVideos(videos)
	} else {
		w.WriteHeader(405)
		fmt.Println(w, "Method not supported!")
	}
}
