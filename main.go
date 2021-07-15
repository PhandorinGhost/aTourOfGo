package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}
type Cache struct {
	sync.Mutex
	db map[string]bool
}

func NewCache() *Cache {
	return &Cache{
		db: make(map[string]bool, 10),
	}
}

func (ch *Cache) get(key string) bool {
	ch.Lock()
	defer ch.Unlock()
	return ch.db[key]
}
func (ch *Cache) set(key string, value bool) {
	ch.Lock()
	defer ch.Unlock()
	ch.db[key] = value
}

var cache = NewCache()
var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	if cache.get(url) {
		fmt.Printf("already checked: %s\n", url)
		return
	}
	fmt.Printf("start crawl: %s\n", url)
	cache.set(url, true)

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
}
func main() {
	wg.Add(1)
	Crawl("https://google.com/", 4, fetcher)
	wg.Wait()
}

type fakeFetcher map[string]*fakeData
type fakeData struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("Not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://google.com/": &fakeData{
		"Hello from Google",
		[]string{
			"https://google.com/img/",
			"https://google.com/cmd/",
		},
	},
	"https://google.com/img/": &fakeData{
		"Packages",
		[]string{
			"https://google.com/",
			"https://google.com/cmd/",
			"https://google.com/img/raw/",
			"https://google.com/img/jpg/",
		},
	},
	"https://google.com/img/raw/": &fakeData{
		"Image raw",
		[]string{
			"https://google.com/",
			"https://google.com/img/",
		},
	},
	"https://google.com/img/jpg/": &fakeData{
		"Image jpg",
		[]string{
			"https://google.com/",
			"https://google.com/img/",
		},
	},
}
