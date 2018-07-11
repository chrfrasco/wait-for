package main

import (
	"errors"
	"net/http"
	"testing"
)

func TestHappyPath(t *testing.T) {
	calls := []string{}

	mockGet := func(url string) (*http.Response, error) {
		calls = append(calls, url)
		return &http.Response{
			StatusCode: 200,
			Status:     "200 OK",
		}, nil
	}

	urlWaiter := newURLWaiter(mockGet)
	urls := []string{"http://google.com"}
	urlWaiter.waitForURLs(urls)

	if !sliceEqual(calls, urls) {
		t.Errorf("expected: %v\n got: %v", urls, calls)
	}
}

func TestBailsAfterError(t *testing.T) {
	calls := []string{}

	mockGet := func(url string) (*http.Response, error) {
		calls = append(calls, url)
		return nil, errors.New("")
	}

	urlWaiter := newURLWaiter(mockGet)
	urls := []string{"http://google.com"}
	urlWaiter.waitForURLs(urls)

	expected := []string{}

	if !sliceEqual(expected, calls) {
		t.Errorf("expected: %v\n got: %v", urls, calls)
	}
}

func sliceEqual(a []string, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
