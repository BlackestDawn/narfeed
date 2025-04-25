package main

const (
	defaultListenPort = "8181"
	defaultTimespan   = "6h"
)

const (
	sectionsUsage = "Comma separated string of sections to read from"
	pagesUsage    = "Number of pages (per section) to read"
	modeUsage     = "Print output to console instead of providing an atom feed"
	helpUsage     = "Print (this) help screen"
	paginateUsage = "Print one post at a time. Only used in console mode"
	portUsage     = "Port for webserver to listen on. Only used in atom mode"
	timespanUsage = "Interval between checking for new posts"
)
