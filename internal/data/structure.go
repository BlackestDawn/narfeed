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
	pageCount       int
	items           map[string]FeedItem
	LastDisplayTime time.Time
}

func NewFeedData(sectionStr string, pageCount int) (*FeedData, error) {
	retVal := new(FeedData)

	retVal.URL = BaseURL
	retVal.sections = validateSections(sectionStr)
	retVal.pageCount = pageCount
	retVal.items = make(map[string]FeedItem)

	return retVal, nil
}
