package main

import "net/url"

func main() {
	var urlStr = "http://127.0.0.1:8102"
	url, err := url.Parse(urlStr)
	if err != nil {
		println(err.Error())
	}
	println(url.String())
	println(url.Scheme)
}
