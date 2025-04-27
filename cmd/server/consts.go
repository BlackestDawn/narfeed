package main

const (
	defaultListenPort = "8181"
	defaultMode       = "atom"
	defaultInterval   = "3h"
)

const (
	usageSections = "Comma separated string of sections to read from"
	usageTags     = "Comma separated string of tags to read from, will skip sections"
	usagePages    = "Number of pages (per section) to read"
	usageMode     = "Set output mode. Currently support atom or console"
	usageHelp     = "Print (this) help screen"
	usagePaginate = "Print one post at a time. Only used in console mode"
	usagePort     = "Port for webserver to listen on. Only used in atom mode"
	usageInterval = "Time interval between checking for updates"
)
