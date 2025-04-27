package data

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (f *FeedData) collectAll() {
	if len(f.tags) == 0 {
		//if len(f.tags) == 0 || f.tags[0] == "" {
		for _, section := range f.sections {
			for i := 1; i <= f.pageCount; i++ {
				var url string
				if section != "inspirational" {
					url = fmt.Sprintf("%s/%s/page/%d", f.URL, section, i)
				} else {
					url = fmt.Sprintf("%s/tag/%s/page/%d", f.URL, section, i)
				}
				f.parsePage(url)
			}
		}
	} else {
		for _, tag := range f.tags {
			for i := 1; i <= f.pageCount; i++ {
				url := fmt.Sprintf("%s/tag/%s/page/%d", f.URL, tag, i)
				f.parsePage(url)
			}
		}
	}
}

func (f *FeedData) parsePage(url string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	doc.Find(".post").Each(func(i int, s *goquery.Selection) {
		headerField := s.Find("div.post_header").Text()
		splitHeader := strings.Split(headerField, "|")
		timeText := strings.TrimSpace(splitHeader[len(splitHeader)-1])
		cTime, err := time.Parse("January _2, 2006", timeText)
		if err != nil {
			log.Printf("Could not parse time from: %s. Using current time instead", timeText)
			cTime = now
		}

		titleField := s.Find("h1.storytitle")
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

		f.add(item)
	})
}
