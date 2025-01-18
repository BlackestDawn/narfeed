package data

import (
	"bufio"
	"fmt"
	"os"
)

func (f *FeedData) PrintToConsole(paginate bool) {
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
