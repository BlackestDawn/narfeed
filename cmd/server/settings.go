package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BlackestDawn/nar-feed/internal/data"
	"github.com/spf13/pflag"
)

type settings struct {
	sections []string
	tags     []string
	pages    int
	paginate bool
	interval time.Duration
	mode     string
	port     string
}

func newSettings() (s *settings, err error) {
	s = new(settings)

	flags := pflag.NewFlagSet("main", pflag.ContinueOnError)
	sections := flags.StringP("sections", "s", data.DefaultSection, usageSections)
	tagsString := flags.StringP("tags", "t", data.DefaultTags, usageTags)
	pages := flags.IntP("pages", "p", data.DefaultPageCount, usagePages)
	paginate := flags.BoolP("paginate", "a", false, usagePaginate)
	interval := flags.StringP("interval", "i", defaultInterval, usageInterval)
	mode := flags.StringP("mode", "m", defaultMode, usageMode)
	port := flags.StringP("listenport", "l", defaultListenPort, usagePort)
	printHelp := flags.BoolP("help", "h", false, usageHelp)

	err = flags.Parse(os.Args[1:])
	if err != nil {
		return
	}

	if *printHelp {
		printHelpText(flags)
		os.Exit(0)
	}

	s.sections = strings.Split(*sections, ",")
	if *tagsString != "" {
		s.tags = strings.Split(*tagsString, ",")
	}
	s.pages = *pages
	s.paginate = *paginate
	s.interval, err = time.ParseDuration(*interval)
	s.mode = *mode
	s.port = *port
	return
}

func printHelpText(flags *pflag.FlagSet) {
	fmt.Println("NAR Feed server")
	fmt.Println()
	flags.PrintDefaults()
}
