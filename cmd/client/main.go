package main

import (
	"fmt"

	flag "github.com/spf13/pflag"

	"github.com/BlackestDawn/nar-feed/internal/data"
)

func main() {
	var sections string
	flag.StringVarP(&sections, "sections", "s", data.DefaultSection, sectionsUsage)

	var pages int
	flag.IntVarP(&pages, "pages", "p", data.DefaultPageCount, pagesUsage)

	var help bool
	flag.BoolVarP(&help, "help", "h", false, helpUsage)

	flag.Parse()

	if help {
		fmt.Println("NAR Feed cmd Client")
		fmt.Println()
		flag.PrintDefaults()
		return
	}

	items, err := data.NewFeedData(sections, pages)
	if err != nil {
		panic(err)
	}
	items.CollectAll()

	items.PrintToConsole(false)
}
