package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type MuxMap struct {
	m   map[string]int
	mux sync.Mutex
	wg  sync.WaitGroup
}

func (mm *MuxMap) do(fn func(m map[string]int) int) int {
	mm.mux.Lock()
	rv := fn(mm.m)
	mm.mux.Unlock()
	return rv
}

func (mm *MuxMap) AddWG() {
	mm.mux.Lock()
	mm.wg.Add(1)
	mm.mux.Unlock()
}

func (mm *MuxMap) DoneWG() {
	mm.mux.Lock()
	mm.wg.Done()
	mm.mux.Unlock()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, mm MuxMap) {
	if depth <= 0 {
		return
	}

	alreadyDone := func(m map[string]int) int {
		v := m[url]
		m[url]++
		// fmt.Printf("%v %v\n", v, url)
		return v
	}

	if mm.do(alreadyDone) > 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, ux := range urls {
		// fmt.Printf("Adding wg for: %v (from: %v)\n", ux, url)
		mm.AddWG()
		go func(u string, d int, f Fetcher, m MuxMap) {
			Crawl(u, d, f, m)
			defer mm.DoneWG()
		}(ux, depth-1, fetcher, mm)
	}
	return
}

func main() {
	mm := MuxMap{m: make(map[string]int)}
	Crawl("http://golang.org/", 4, fetcher, mm)
	time.Sleep(time.Second) // THIS IS SO LAME!
	mm.wg.Wait()            // Why doesn't this work as expected?
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
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
