package main

import (
	"github.com/antonholmquist/jason"
	"os"
	"os/exec"
)

var Npm = func(dir string) string{

	err := os.Chdir(dir)
	if err != nil {
		println(err.Error())
		return ""
	}

	cmd := exec.Command("npm", "list", "--json=true")
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return ""
	}

	debug(stdout)

	v, err := jason.NewObjectFromBytes(stdout)
	// name, err := v.GetString("name")

	dependencies, err := v.GetObject("dependencies")

	// todo : 配列への対応
	var keyword string
	for name, _ := range dependencies.Map() {
		keyword = name
	}


	return keyword
}
