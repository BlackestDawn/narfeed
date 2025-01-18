package main

import (
	"log"
)

func main() {
	serverConf, err := newServerConf()
	if err != nil {
		log.Fatal(err)
	}

	serverConf.run()
}
