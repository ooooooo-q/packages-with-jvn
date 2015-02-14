package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandNpm,
}

var commandNpm = cli.Command{
	Name:  "npm",
	Usage: "",
	Description: `
`,
	Action: doNpm,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doNpm(c *cli.Context) {
	dir := c.Args()[0]
	keywords := Npm(dir)
	println(keywords)
}
