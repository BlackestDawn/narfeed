package main

import (
	"fmt"
	"os"
)

func (c *serverConf) printHelp(eCode int) {
	fmt.Println("NAR Feed Server")
	fmt.Println()
	fmt.Println("Usage: server [mode] <options...>")
	fmt.Println()
	fmt.Println("Currently supported modes are 'atom' and 'console'. It also accepts 'help' to print this help text")
	fmt.Println("  In atom mode it will act as a simple webserver providing an atom feed.")
	fmt.Println("  In console mode it will output to the console/terminal")
	fmt.Println()
	fmt.Println("Options:")
	c.flagSet.PrintDefaults()
	os.Exit(eCode)
}
