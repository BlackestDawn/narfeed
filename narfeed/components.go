package narfeed

import (
	"fmt"
	"time"
)

func (f *FeedData) add(item FeedItem) {
	f.items[item.ID] = item
}

func (f *FeedData) GetIDs() (retVal []string) {
	for key := range f.items {
		retVal = append(retVal, key)
	}

	return
}

func (f *FeedData) GetItem(id string) (FeedItem, error) {
	if item, ok := f.items[id]; ok {
		return item, nil
	}
	return FeedItem{}, fmt.Errorf("could not find item with ID: %s", id)
}

func (f *FeedData) getAllItems() (returnItems []FeedItem) {
	now := time.Now()

	for _, item := range f.items {
		returnItems = append(returnItems, item)
	}

	f.LastDisplayTime = now

	return
}

func (f *FeedData) GetNewItems() (returnItems []FeedItem) {
	f.collectAll()
	now := time.Now()

	for _, item := range f.items {
		if item.CreatedAt.After(f.LastDisplayTime) {
			returnItems = append(returnItems, item)
		}
	}

	f.LastDisplayTime = now

	return
}
