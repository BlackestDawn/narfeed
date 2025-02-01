package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BlackestDawn/nar-feed/internal/data"
	"github.com/spf13/pflag"
)

func parseArgs() *serverConf {
	conf := new(serverConf)

	defaultParsedTimespan, err := time.ParseDuration(defaultTimespan)
	if err != nil {
		log.Fatal(err)
	}

	conf.flagSet = pflag.NewFlagSet("common", pflag.ContinueOnError)
	sections := conf.flagSet.StringP("sections", "s", data.DefaultSection, sectionsUsage)
	pages := conf.flagSet.IntP("pages", "p", data.DefaultPageCount, pagesUsage)
	paginate := conf.flagSet.BoolP("paginate", "a", false, paginateUsage)
	port := conf.flagSet.StringP("listenport", "l", defaultListenPort, portUsage)
	timespan := conf.flagSet.DurationP("timespan", "t", defaultParsedTimespan, timespanUsage)
	help := conf.flagSet.BoolP("help", "h", false, helpUsage)

	err = conf.flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		conf.printHelp(1)
	}

	conf.mode = os.Args[1]
	conf.sections = *sections
	conf.pages = *pages
	conf.paginate = *paginate
	conf.port = *port
	conf.help = *help
	conf.interval = timespan.Abs()

	if conf.mode == "help" || conf.help {
		conf.printHelp(0)
	}

	return conf
}
