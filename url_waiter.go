package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type getRequestFunc = func(url string) (resp *http.Response, err error)

type urlWaiter struct {
	wg  *sync.WaitGroup
	get getRequestFunc
}

func newURLWaiter(get getRequestFunc) *urlWaiter {
	uw := new(urlWaiter)
	uw.get = get
	return uw
}

func (uw *urlWaiter) waitForURLs(urls []string) {
	uw.wg = new(sync.WaitGroup)
	uw.wg.Add(len(urls))

	fmt.Printf("waiting for %d urls...\n", len(urls))

	for _, url := range urls {
		go uw.waitForURL(url)
	}

	uw.wg.Wait()
}

func (uw *urlWaiter) waitForURL(url string) {
	defer uw.wg.Done()

	url = normalize(url)

	for {
		resp, err := uw.get(url)
		if err != nil {
			log.Fatalf("could not fetch %s: %v", url, err)
		}

		if 200 <= resp.StatusCode && resp.StatusCode <= 299 {
			fmt.Printf("%s is up: %s\n", url, resp.Status)
			break
		}

		time.Sleep(time.Second)
	}

}

func normalize(url string) string {
	if strings.HasPrefix(url, ":") {
		url = fmt.Sprintf("http://localhost%s", url)
	}

	return url
}
