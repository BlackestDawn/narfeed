package data

import (
	"log"
)

const (
	BaseURL          = "https://notalwaysright.com"
	DefaultPageCount = 4
	DefaultSection   = "newest"
	FeedID           = "91f04a6d-ff1c-4d86-8d50-9962f640eda7"
	FeedTitle        = "Not Always Right"
	FeedDescription  = ""
)

type SectionType string

func getValidSections() map[SectionType]struct{} {
	return map[SectionType]struct{}{
		"newest":        {},
		"right":         {},
		"working":       {},
		"romantic":      {},
		"related":       {},
		"learning":      {},
		"friendly":      {},
		"healthy":       {},
		"legal":         {},
		"inspirational": {},
	}
}

// takes a comma-separated list of sections and returns a slice of only the valid ones
func validateSections(sections []string) []string {
	validSect := getValidSections()
	var validated []string

	for _, val := range sections {
		if _, ok := validSect[SectionType(val)]; ok {
			validated = append(validated, val)
		} else {
			log.Println("unknown section:", val)
		}
	}

	return validated
}
