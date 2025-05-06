package main

import (
	"log"

	"github.com/BlackestDawn/narfeed"
)

func main() {
	settings, err := newSettings()
	if err != nil {
		log.Fatal(err)
	}

	feed, err := narfeed.NewFeedData(settings.sections, settings.tags, settings.pages)
	if err != nil {
		log.Fatal(err)
	}

	switch settings.mode {
	case "console":
		runConsoleServer(feed, settings)
	case "atom":
		runWebServer(feed, settings)
	default:
		log.Fatal("unknown mode:", settings.mode)
	}
}
