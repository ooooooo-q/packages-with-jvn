package main

import (
	"io/ioutil"
	"net/http"
)

var Jvn = func(keyword string) {

	url := "http://jvndb.jvn.jp/myjvn?method=getVulnOverviewList"
	url += "&rangeDatePublic=m"
	url += "&keyword=" + keyword

	debug(url)

	response, err := http.Get(url)
	if err != nil {
		println(err.Error())
		return
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println(err.Error())
			return
		}
		println(string(contents))
	}
}
