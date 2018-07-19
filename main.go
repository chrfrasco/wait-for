package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	openAfterDone := flag.Bool("open", false, "Open the URL after it's up")
	flag.Parse()

	urls := flag.Args()
	waitForUrls(urls)

  // TODO: should not have to wait for ALL urls
	if *openAfterDone {
		for _, url := range urls {
			openURL(url)
		}
	}
}

func waitForUrls(urls []string) {
	urlWaiter := newURLWaiter(http.Get)
	urlWaiter.waitForURLs(urls)
}

func openURL(url string) {
	cmd := exec.Command("open", normalize(url))
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
