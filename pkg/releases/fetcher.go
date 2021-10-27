package releases

import (
	"fmt"
	"github.com/gocolly/colly"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

type elementMapper = func(element *colly.HTMLElement) Manga

type ReleaseFetcher interface {
	Fetch() ([]Manga, error)
}

type releaseFetcher struct {
	http          *http.Transport
	url           string
	querySelector string
	mapper        elementMapper
}

func (r releaseFetcher) Fetch() ([]Manga, error) {
	var releases []Manga

	collector := colly.NewCollector()
	collector.WithTransport(r.http)

	collector.OnHTML(r.querySelector, func(element *colly.HTMLElement) {
		releases = append(releases, r.mapper(element))
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	pwd, _ := os.Getwd()
	filePath := "file://" + path.Join(pwd, r.url)
	if err := collector.Visit(filePath); err != nil {
		return nil, err
	}

	return releases, nil
}

func NewReleaseFetcher(querySelector string, url string, mapper elementMapper) ReleaseFetcher {
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	return releaseFetcher{
		http:          t,
		querySelector: querySelector,
		url:           url,
		mapper:        mapper,
	}
}
