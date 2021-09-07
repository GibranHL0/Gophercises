package site

import "strings"

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func filter(links []string, keepFn func(string) bool) []string {
	var filtered []string

	for _, link := range links {
		if keepFn(link) {
			filtered = append(filtered, link)
		}
	}

	return filtered
}

func ConvertToUrls(links []string) []Url {
	urls := make([]Url, 0, len(links))

	for _, link := range links {
		urls = append(urls, Url{Loc: link})
	}

	return urls
}

func Bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})

	var queue map[string]struct{}

	nextQueue := map[string]struct{}{
		urlStr: {},
	}

	for i := 0; i <= maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]struct{})

		for url := range queue {			if _, ok := seen[url]; ok {
				continue
			}

			seen[url] = struct{}{}

			for _, link := range getPages(urlStr){
				nextQueue[link] = struct{}{}
			}
		}
	}

	allLinks := make([]string, 0, len(seen))

	for url := range seen {
		allLinks = append(allLinks, url)
	}

	return allLinks
}