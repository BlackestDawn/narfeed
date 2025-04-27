package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BlackestDawn/nar-feed/internal/data"
)

func runWebServer(feed *data.FeedData, settings *settings) {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomFeed := feed.GenerateAtomFeed()

		atom, err := atomFeed.ToAtom()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, atom)
	})

	server := &http.Server{
		Handler: serverMux,
		Addr:    ":" + settings.port,
	}
	log.Println("starting webserver...")
	log.Print(server.ListenAndServe())

}
