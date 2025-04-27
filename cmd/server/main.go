package main

import (
	"log"

	"github.com/BlackestDawn/nar-feed/internal/data"
)

func main() {
	settings, err := newSettings()
	if err != nil {
		log.Fatal(err)
	}

	feed, err := data.NewFeedData(settings.sections, settings.tags, settings.pages)
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
