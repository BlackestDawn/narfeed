package data

import (
	"time"

	"github.com/gorilla/feeds"
)

func (f *FeedData) GenerateAtomFeed() *feeds.Feed {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       FeedTitle,
		Link:        &feeds.Link{Href: BaseURL},
		Description: FeedDescription,
		Created:     now,
	}

	for _, item := range f.GetAllItems() {
		feed.Add(&feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.URL},
			Description: item.Content,
			Created:     item.CreatedAt,
			Updated:     item.UpdatedAt,
			Id:          item.ID,
		})
	}

	return feed
}
