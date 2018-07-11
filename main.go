package main

import (
	"os"
)

func main() {
	urls := os.Args[1:]
	waitForUrls(urls)
}

func waitForUrls(urls []string) {
	urlWaiter := new(urlWaiter)
	urlWaiter.waitForURLs(urls)
}
