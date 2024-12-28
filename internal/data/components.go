package data

import "fmt"

func (f *FeedData) Add(item FeedItem) {
	if _, ok := f.items[item.ID]; ok {
		return
	}

	f.items[item.ID] = item
}

func (f *FeedData) GetIDs() []string {
	var retVal []string

	for key := range f.items {
		retVal = append(retVal, key)
	}

	return retVal
}

func (f *FeedData) GetItem(id string) (FeedItem, error) {
	if item, ok := f.items[id]; ok {
		return item, nil
	}
	return FeedItem{}, fmt.Errorf("Could no find item with ID: %s", id)
}

func (f *FeedData) GetAllItems() []FeedItem {
	var retVal []FeedItem

	for _, item := range f.items {
		retVal = append(retVal, item)
	}

	return retVal
}
