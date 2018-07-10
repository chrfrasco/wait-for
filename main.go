package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	urls := os.Args[1:]

	var wg sync.WaitGroup
	wg.Add(len(urls))

	fmt.Printf("waiting for %d urls...\n", len(urls))

	for _, url := range urls {
		go waitFor(url, &wg)
	}

	wg.Wait()
}

func waitFor(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	url = normalize(url)

	for {
		res, err := http.Get(url)
		if err != nil {
			continue
		}

		if res != nil {
			time.Sleep(time.Millisecond * 100)
			break
		}
	}

	fmt.Printf("%s is up\n", url)
}

func normalize(url string) string {
	if strings.HasPrefix(url, ":") {
		url = fmt.Sprintf("http://localhost%s", url)
	}

	return url
}
