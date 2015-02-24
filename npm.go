package main

import (
	"github.com/antonholmquist/jason"
	"os"
	"os/exec"
)

var Npm = func(dir string) []string {

	err := os.Chdir(dir)
	if err != nil {
		println(err.Error())
		return nil
	}

	cmd := exec.Command("npm", "list", "--json=true")
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return nil
	}

	v, err := jason.NewObjectFromBytes(stdout)
	if err != nil {
		println(err.Error())
		return nil
	}

	dependencies, err := v.GetObject("dependencies")
	if err != nil {
		println(err.Error())
		return nil
	}

	var keywords []string
	for name, _ := range dependencies.Map() {
		keywords = append(keywords, name)
	}

	return keywords
}
