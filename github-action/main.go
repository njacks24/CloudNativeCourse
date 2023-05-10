package main

import (
	"log"
	"github.com/njacks24/github-action/microservice"


func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
