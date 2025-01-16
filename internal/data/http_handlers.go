package data

import (
	"log"
	"net/http"
)

func (f *FeedData) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	f.CollectAll()
	atomFeed := f.GenerateAtomFeed()

	atom, err := atomFeed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(atom))
}
