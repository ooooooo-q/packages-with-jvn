package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandCheck,
}

var commandCheck = cli.Command{
	Name:  "check",
	Usage: "",
	Description: `
`,
	Action: doCheck,
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

func doCheck(c *cli.Context) {
}
