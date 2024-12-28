package data

import "fmt"

func (f *FeedData) PrintToConsole(paginate bool) {
	for _, item := range f.items {
		fmt.Println(item.Title)
		fmt.Println()
		fmt.Println(item.Content)
		fmt.Printf("\n\n")
		if paginate {
			fmt.Println("press enter key to continue...")
			fmt.Scan()
		}
		fmt.Println()
	}
}
