package main

import (
	"fmt"
	"os"

	"github.com/BlackestDawn/nar-feed/internal/data"
)

func main() {
	settings, err := newSettings()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	feed, err := data.NewFeedData(settings.sections, settings.tags, settings.pages)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	feed.PrintToConsole(settings.paginate)
}
