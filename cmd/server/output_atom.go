package main

import (
	"log"
	"net/http"
)

func (c *serverConf) outputAtom() {
	http.HandleFunc("/", c.feedData.HandleHttpRequest)

	log.Println("Started on port:", c.port)
	err := http.ListenAndServe(":"+c.port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
