package main

import (
	"log"
	"net/http"
)

var members []Member

func main() {
	members = initMembers()
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
