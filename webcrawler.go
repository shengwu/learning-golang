package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, bodyc chan string, urlc chan []string, errc chan error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher,
	visited map[string]bool) {
	if depth <= 0 {
		return
	}
	if !visited[url] {
		visited[url] = true
		bodyc := make(chan string)
		urlsc := make(chan []string)
		errc := make(chan error)
		go fetcher.Fetch(url, bodyc, urlsc, errc)
		select {
		case body := <-bodyc:
			urls := <-urlsc
			fmt.Printf("found: %s %q\n", url, body)
			for _, u := range urls {
				Crawl(u, depth-1, fetcher, visited)
			}
		case err := <-errc:
			fmt.Println(err)
			return
		}
	}
	return
}

func main() {
	visited := make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher, visited)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, bodyc chan string,
	urlc chan []string, errc chan error) {
	if res, ok := f[url]; ok {
		bodyc <- res.body
		urlc <- res.urls
		errc <- nil
	}
	bodyc <- ""
	urlc <- nil
	errc <- fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
