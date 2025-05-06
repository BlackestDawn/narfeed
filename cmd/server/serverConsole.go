package main

import (
	"log"
	"time"

	"github.com/BlackestDawn/narfeed/narfeed"
	"github.com/kardianos/service"
)

//var logger service.Logger

type program struct {
	feed     *narfeed.FeedData
	settings *settings
}

func (p *program) Start(s service.Service) (err error) {
	go p.run()
	return
}

func (p *program) Stop(s service.Service) (err error) {
	return
}

func (p *program) run() {
	p.feed.PrintToConsole(p.settings.paginate)
	time.Sleep(p.settings.interval)
}

func runConsoleServer(feed *narfeed.FeedData, settings *settings) {
	log.Println("starting console server...")

	svcConfig := &service.Config{
		Name:        "NARFeedServer",
		DisplayName: "NAR Feed Server",
		Description: "Console server for collecting and siplaying posts on NotAlwaysRight",
	}

	prg := &program{
		feed:     feed,
		settings: settings,
	}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatal(logger.Error(err))
	}
}
