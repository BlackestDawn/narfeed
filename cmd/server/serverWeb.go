package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BlackestDawn/narfeed"
)

func runWebServer(feed *narfeed.FeedData, settings *settings) {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atom, err := feed.GenerateAtomFeed()
		if err != nil {
			log.Fatalf("GenerateAtomFeed: %v", err)
		}

		fmt.Fprint(w, atom)
	})

	server := &http.Server{
		Handler:           serverMux,
		Addr:              ":" + settings.port,
		ReadHeaderTimeout: time.Second * 20,
	}
	log.Println("starting webserver...")
	log.Print(server.ListenAndServe())

}
