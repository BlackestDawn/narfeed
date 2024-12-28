package generateoutput

import (
	"fmt"

	"github.com/BlackestDawn/nar-feed/internal/collect"
)

func PrintConsole(items []collect.FeedItem) {
	for _, item := range items {
		fmt.Println(item.Title)
		fmt.Println()
		fmt.Println(item.Content)
		fmt.Printf("\n\n\n")
	}
}
