package main

import (
	"SortShapesService/myroutes"
	"log"
	"net/http"
)

func main() {

	router := myroutes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
