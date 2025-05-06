package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BlackestDawn/narfeed"
	"github.com/spf13/pflag"
)

type settings struct {
	sections []string
	tags     []string
	pages    int
	paginate bool
}

func newSettings() (s *settings, err error) {
	s = new(settings)

	flags := pflag.NewFlagSet("main", pflag.ContinueOnError)

	sectionsString := flags.StringP("sections", "s", narfeed.DefaultSection, usageSections)
	tagsString := flags.StringP("tags", "t", narfeed.DefaultTags, usageTags)
	pages := flags.IntP("pages", "p", narfeed.DefaultPageCount, usagePages)
	paginate := flags.BoolP("paginate", "a", false, usagePaginate)
	printHelp := flags.BoolP("help", "h", false, usageHelp)

	err = flags.Parse(os.Args[1:])
	if err != nil {
		printHelpText(flags)
		return
	}

	if *printHelp {
		printHelpText(flags)
		os.Exit(0)
	}

	s.sections = strings.Split(*sectionsString, ",")
	if *tagsString != "" {
		s.tags = strings.Split(*tagsString, ",")
	}
	s.pages = *pages
	s.paginate = *paginate
	return
}

func printHelpText(flags *pflag.FlagSet) {
	fmt.Println("NAR Feed cli Client")
	fmt.Println()
	flags.PrintDefaults()
}
