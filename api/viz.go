package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/gocolly/colly"
)

type Manga struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

var location, _ = time.LoadLocation("UTC")
var year, month, day = time.Now().In(location).Date()

func CollectVizReleases() []Manga {
	var allVizReleases []Manga

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	collector := colly.NewCollector()
	collector.WithTransport(t)

	collector.OnHTML(".manga-books article", func(element *colly.HTMLElement) {
		temp := Manga{}
		temp.Name = element.ChildText(".color-off-black")
		temp.Image = element.ChildAttr("a.product-thumb img", "data-original")
		temp.Link = "https://viz.com" + element.ChildAttr("a.product-thumb", "href")
		allVizReleases = append(allVizReleases, temp)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	//pwd, _ := os.Getwd()
	s := fmt.Sprintf("pages/viz-%d-%d-%d.html", int(year), int(month), int(day))
	_, currentFilePath, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFilePath)
	//s := fmt.Sprintf("pages/ci.yml")
	collector.Visit("file://" + path.Join(dir, s))

	return allVizReleases
}

func Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CollectVizReleases())
}
