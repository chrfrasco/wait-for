package main

import (
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	waitForUrls(urls)
}

func waitForUrls(urls []string) {
	urlWaiter := newURLWaiter(http.Get)
	urlWaiter.waitForURLs(urls)
}
