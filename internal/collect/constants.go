package collect

import "time"

const BaseURL = "https://notalwaysright.com"
const DefaultPageCount = 3
const DefaultSection = "newest"
const FeedID = "91f04a6d-ff1c-4d86-8d50-9962f640eda7"

type FeedItem struct {
	Title     string
	URL       string
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
}
