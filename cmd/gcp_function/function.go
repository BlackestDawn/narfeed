package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BlackestDawn/nar-feed/internal/data"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("NAR-feed", Process)
}

func Process(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(0)
	logger := log.New(funcframework.LogWriter(r.Context()), "", 0)

	settings := setting{
		Sections: []string{data.DefaultSection},
		Pages:    data.DefaultPageCount,
	}

	err := r.ParseForm()
	if err != nil {
		logger.Fatalf("ParseForm: %v", err)
	}

	if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
		logger.Fatalf("json decoding: %v", err)
	}

	feed, err := data.NewFeedData(settings.Sections, settings.Tags, settings.Pages)
	if err != nil {
		logger.Fatalf("NewFeedData: %v", err)
	}

	atomFeed, err := feed.GenerateAtomFeed()
	if err != nil {
		logger.Fatalf("GenerateAtomFeed: %v", err)
	}

	fmt.Fprint(w, atomFeed)
}
