package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/darthxibalba/learning-go/html-parse"
)

const xmlns = "https://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:loc`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "the url that will be used to generate a sitemap")
	maxDepth := flag.Int("depth", 3, "the maximum number of link tree depth")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}

// Hyper-Link cases
/*
  /some-path
  https://gophercises.com/some-path
  #fragment
  /some-path#fragment (don't care)
  https://gophercises.com/some-path#fragment (don't care)
  mailto:jon@calhoun.io (don't care)
  Possibly other paths:
  //gophercises.com/some.css (css for not caring if http/https)
*/

/*
  R1. GET the webpage
  R2. Parse all the links on the page
  R3. Build proper urls with our links
  R4. Filter out any links w/ a diff domain
  R5. Find all pages (BFS)
  6.  Print out data as XML
*/

type empty struct{}

func bfs(urlStr string, maxDepth int) []string {
	// Map to empty struct since it takes less memory compared to bools
	seen := make(map[string]empty)
	var queue map[string]empty
	nextQueue := map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i < maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]empty)

		for url, _ := range queue {
			if _, ok := seen[url]; ok { // if key exists in map
				continue
			}
			// Add url value to seen if we haven't seen it
			seen[url] = empty{}
			// Range over all links and for each link, put them in next queue
			for _, link := range get(url) {
				nextQueue[link] = empty{}
			}
		}
	}
	// Parse back into slice
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//io.Copy(os.Stdout, resp.Body)

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(htmlReader io.Reader, base string) []string {
	links, _ := link.Parse(htmlReader)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		// https://gophercises.com
		// https://gophercises.com/some-path
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
