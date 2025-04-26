package data

import (
	"time"
)

type FeedItem struct {
	Title     string
	URL       string
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
}

type FeedData struct {
	URL             string
	sections        []string
	tags            []string
	pageCount       int
	items           map[string]FeedItem
	LastDisplayTime time.Time
}

func NewFeedData(sections []string, tags []string, pageCount int) (feedData *FeedData, err error) {
	feedData = new(FeedData)

	feedData.URL = BaseURL
	feedData.sections = sections
	feedData.tags = tags
	feedData.pageCount = pageCount
	feedData.items = make(map[string]FeedItem)

	return
}
