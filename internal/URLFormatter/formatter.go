package urlformatter

import (
	"log/slog"
	"net/url"
	"regexp"
	"strings"
)

var Logger slog.Logger

func destructPath(u string) (string, int) {
	length, num := len(u), 0

	if length < 4 {
		return "", 1
	}

	found := true

	for found {
		u, found = strings.CutPrefix(u, "../")
		num++
	}

	return u, num
}

func format(originalURL url.URL, u string) (url.URL, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		Logger.Error("Error parsing URL")
		return url.URL{}, err
	}
	matched, err := regexp.MatchString(`^https?://www\.[a-zA-Z0-9]+\.[a-zA-Z]{2,6}/?.*$`, u)
	if matched {
		return *parsedURL, nil
	}

	matched, err = regexp.MatchString(`^//www\.[a-zA-Z0-9]+\.[a-zA-Z0-9]{2,6}/?.*$`, u)
	if matched {
		parsedURL.Scheme = originalURL.Scheme
		return *parsedURL, nil
	}
	matched, err = regexp.MatchString(`^[a-zA-Z-0-9.]+`, u)
	if matched {
		originalURL.JoinPath(u)
		return originalURL, nil

	}
	matched, err = regexp.MatchString(`^\./[a-zA-Z-0-9.]+`, u)
	if matched {
		originalURL.JoinPath(u[2:])
		return originalURL, nil

	}
	matched, err = regexp.MatchString(`^[a-zA-Z0-9]+/.*`, u)
	if matched {
		originalURL.JoinPath(u)
		return originalURL, nil
	}
	matched, err = regexp.MatchString(`^(\.\./)+[a-zA-Z0-9.@?]+`, u)
	if matched {
		path, traversed := destructPath(u)
		pathWords := strings.Split(originalURL.Path, "/")
		pathLen := len(pathWords)
		finalLen := pathLen - traversed
		finalPath := ""

		for index := range finalLen {
			finalPath += pathWords[index]
		}

		finalPath += path
		originalURL.Path = finalPath

		return originalURL, nil
	}

	if err != nil {
		Logger.Error("Error matching URL")
		return url.URL{}, err
	}
	return url.URL{}, nil
}

func Traverse(originalURL url.URL, urls []string) []string {
	formattedURLs := make(map[string]bool)

	for _, u := range urls {
		newUrl, err := format(originalURL, u)
		if err != nil {
			continue
		}
		newUrlStr := newUrl.String()
		if len(newUrlStr) > 0 && !formattedURLs[newUrlStr] {
			formattedURLs[newUrlStr] = true
		}
	}

	array := []string{}
	for u := range formattedURLs {
		array = append(array, u)
	}
	return array
}
