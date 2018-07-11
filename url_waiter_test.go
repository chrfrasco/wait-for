package main

import (
	"net/http"
	"testing"
)

func makeMockGet(fn func(url string)) func(url string) (*http.Response, error) {
	return func(url string) (*http.Response, error) {
		fn(url)
		return &http.Response{}, nil
	}
}

func TestURLWaiterBasic(t *testing.T) {
	calls := []string{}
	mockGet := makeMockGet(func(url string) {
		calls = append(calls, url)
	})
	urlWaiter := newURLWaiter(mockGet)
	urlWaiter.waitForURLs([]string{"http://google.com"})
}
