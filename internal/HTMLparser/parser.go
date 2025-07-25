package htmlparser

import (
	"log/slog"
	"net/http"
	"net/url"
	urlformatter "web-crawler/internal/URLFormatter"

	"golang.org/x/net/html"
)

var Logger slog.Logger

func HTMLParser(used *map[string]bool, output *map[string][]string, u string) {

	resp, err := http.Get(u)
	if err != nil {
		Logger.Error("Invalid URL")
		return
	}

	defer resp.Body.Close()
	data, err := html.Parse(resp.Body)
	if err != nil {
		Logger.Error(`Error while reading URL\'s body`)
		return
	}

	var urls []string

	traverse(used, output, u, &urls, data)
	urls = []string{}

	for _, suburls := range *output {
		for _, suburl := range suburls {
			HTMLParser(used, output, suburl)
		}
	}
}

func traverse(used *map[string]bool, output *map[string][]string, u string, urls *[]string, node *html.Node) {
	if (*used)[u] {
		return
	}

	(*used)[u] = true

	if node.Type == html.ElementNode {
		for _, e := range node.Attr {
			if e.Key == "href" {
				*urls = append(*urls, e.Val)
			}
		}
	}

	originalURL, err := url.Parse(u)
	if err != nil {
		Logger.Error("Error while parsing URL")
		return
	}
	formattedURLS := urlformatter.Traverse(*originalURL, *urls)
	(*output)[u] = formattedURLS

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverse(used, output, u, urls, child)
	}
}
