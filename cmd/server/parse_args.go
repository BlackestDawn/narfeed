package main

import (
	"fmt"
	"os"

	"github.com/BlackestDawn/nar-feed/internal/data"
	"github.com/spf13/pflag"
)

func parseArgs() *serverConf {
	conf := new(serverConf)

	conf.flagSet = pflag.NewFlagSet("common", pflag.ContinueOnError)
	sections := conf.flagSet.StringP("sections", "s", data.DefaultSection, sectionsUsage)
	pages := conf.flagSet.IntP("pages", "p", data.DefaultPageCount, pagesUsage)
	paginate := conf.flagSet.BoolP("paginate", "a", false, paginateUsage)
	port := conf.flagSet.StringP("listenport", "l", defaultListenPort, portUsage)
	help := conf.flagSet.BoolP("help", "h", false, helpUsage)

	err := conf.flagSet.Parse(os.Args[2:])
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

	if conf.mode == "help" || conf.mode == "-h" || conf.help {
		conf.printHelp(0)
	}

	return conf
}
