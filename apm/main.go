package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)
	r.Use(apmgorilla.Middleware())
	log.Fatal(http.ListenAndServe(":8000", r))
}
