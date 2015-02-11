package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
	"os/exec"
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

	cmd := exec.Command("npm", "list", "--json=true")
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}
