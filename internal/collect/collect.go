package collect

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func StartCollect() []FeedItem {
	var retVal []FeedItem

	pageCount := DefaultPageCount
	for i := 1; i <= pageCount; i++ {
		url := fmt.Sprintf("%s/%s/page/%d", BaseURL, DefaultSection, i)
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		items := parsePage(res.Body)
		retVal = append(retVal, items...)
	}

	return retVal
}

func parsePage(reader io.Reader) []FeedItem {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	var retVal []FeedItem
	now := time.Now()

	doc.Find(".post").Each(func(i int, s *goquery.Selection) {
		headerField := s.Find("div.post_header").Text()
		titleField := s.Find("h1.storytitle")

		splitHeader := strings.Split(headerField, "|")
		timeText := strings.TrimSpace(splitHeader[len(splitHeader)-1])
		cTime, err := time.Parse("January _2, 2006", timeText)
		if err != nil {
			log.Printf("Could not parse time from: %s. Using current time instead", timeText)
			cTime = now
		}

		link, _ := titleField.Find("a").Attr("href")
		splitLink := strings.Split(strings.TrimRight(link, "/"), "/")
		id := splitLink[len(splitLink)-1]

		item := FeedItem{
			Title:     strings.TrimSpace(titleField.Text()),
			URL:       link,
			ID:        id,
			CreatedAt: cTime,
			UpdatedAt: now,
			Content:   strings.TrimSpace(s.Find("div.storycontent").Text()),
		}

		retVal = append(retVal, item)
	})

	return retVal
}
