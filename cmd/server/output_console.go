package main

func (c *serverConf) outputConsole() {
	c.feedData.CollectAll()

	c.feedData.PrintToConsole(c.paginate)
}
