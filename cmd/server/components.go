package main

import (
	"fmt"
)

func (c *serverConf) run() {
	switch c.mode {
	case "atom":
		c.outputAtom()
	case "console":
		c.outputConsole()
	default:
		fmt.Printf("Unknown server mode: %s\n", c.mode)
		c.printHelp(1)
	}

}
