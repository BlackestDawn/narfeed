package data

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gorilla/feeds"
)

func (f *FeedData) PrintToConsole(paginate bool) {
	f.collectAll()
	for _, item := range f.items {
		fmt.Println(item.Title)
		fmt.Println()
		fmt.Println(item.Content)
		fmt.Printf("\n\n")
		if paginate {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("press enter key to continue...")
			reader.ReadByte()
		}
		fmt.Println()
	}
}

func (f *FeedData) GenerateAtomFeed() (feed *feeds.Feed) {
	f.collectAll()
	now := time.Now()
	feed = &feeds.Feed{
		Title:       FeedTitle,
		Link:        &feeds.Link{Href: BaseURL},
		Description: FeedDescription,
		Created:     now,
	}

	for _, item := range f.getAllItems() {
		feed.Add(&feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.URL},
			Description: item.Content,
			Created:     item.CreatedAt,
			Updated:     item.UpdatedAt,
			Id:          item.ID,
		})
	}

	return
}
