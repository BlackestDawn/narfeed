package main

import (
	"time"

	"github.com/BlackestDawn/nar-feed/internal/data"
	"github.com/spf13/pflag"
)

type serverConf struct {
	mode     string
	sections string
	pages    int
	paginate bool
	port     string
	help     bool
	flagSet  *pflag.FlagSet
	feedData *data.FeedData
	interval time.Duration
}

func newServerConf() (*serverConf, error) {
	conf := parseArgs()

	fd, err := data.NewFeedData(conf.sections, conf.pages)
	if err != nil {
		return &serverConf{}, err
	}

	conf.feedData = fd

	return conf, nil
}
