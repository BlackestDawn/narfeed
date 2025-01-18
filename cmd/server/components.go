package main

import "log"

func (c *serverConf) run() {
	switch c.mode {
	case "atom":
		c.outputAtom()
	case "console":
		c.outputConsole()
	default:
		log.Fatalf("Unknown server mode: %s", c.mode)
	}

}
