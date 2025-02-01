package main

import (
	"fmt"
	"os"
)

func (c *serverConf) printHelp(eCode int) {
	fmt.Println(`NAR Feed Server

Usage: server [mode] <options...>

Currently supported modes are 'atom' and 'console'. It also accepts 'help' to print this help text
  In atom mode it will act as a simple webserver providing an atom feed.
  In console mode it will output to the console/terminal

Options:`)
	c.flagSet.PrintDefaults()
	os.Exit(eCode)
}
