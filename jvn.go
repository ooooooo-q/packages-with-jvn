package main

import (
	"io/ioutil"
	"net/http"
)

var Jvn = func(keyword string) {

	url := "http://jvndb.jvn.jp/myjvn?method=getVulnOverviewList"
	url += "&rangeDatePublic=m"
	url += "&keyword=" + keyword

	xml_str, err := request(url)
	if err != nil {
		println(err.Error())
		return
	}

	println(xml_str)
}

func request(url string) (result string, err error) {

	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}
