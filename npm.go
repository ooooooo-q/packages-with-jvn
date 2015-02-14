package main

import (
	"os"
	"os/exec"
)

var Npm = func(dir string) {

	err := os.Chdir(dir)
	if err != nil {
		println(err.Error())
		return
	}

	cmd := exec.Command("npm", "list", "--json=true")
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}
