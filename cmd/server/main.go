package main

import (
	"log"
	"net/http"

	"github.com/BlackestDawn/nar-feed/internal/data"
)

func main() {
	feedItems, err := data.NewFeedData(data.DefaultSection, data.DefaultPageCount)
	if err != nil {
		panic(err)
	}

	portNum := defaultListenPort

	http.HandleFunc("/", feedItems.HandleHttpRequest)

	log.Println("Started on port", portNum)
	err = http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
