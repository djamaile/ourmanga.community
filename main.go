package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"github.com/gocolly/colly"
)

type Manga struct {
	Name  string
	Image string
	Link  string
}
var year, month, day = time.Now().Date()

func collectYenPressReleases() {
	var allYenReleases []Manga

	collector := colly.NewCollector()

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	collector.WithTransport(t)

	collector.OnHTML(".book-shelf-title-grid", func(element *colly.HTMLElement) {
		temp := Manga{}
		temp.Name = element.ChildText(".book-detail p")
		temp.Image = element.ChildAttr("img", "src")
		temp.Link = element.ChildAttr(".book-detail-links a", "href")
		allYenReleases = append(allYenReleases, temp)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	pwd, _ := os.Getwd()
	collector.Visit("file://" + path.Join(pwd, "pages/yen.html"))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allYenReleases)
}

func collectSevenSeasReleases() {
	var allSevenSeasReleases []Manga

	collector := colly.NewCollector(
		colly.AllowedDomains("sevenseasentertainment.com", "www.sevenseasentertainment.com"),
	)

	collector.OnHTML("div[style='float: left; margin: 0 3px 10px 6px; width: 134px; height: 189px; background: #CECECE;']", func(element *colly.HTMLElement) {
		temp := Manga{}
		temp.Name = element.ChildAttr("a", "title")
		temp.Image = element.ChildAttr("img", "src")
		temp.Link = element.ChildAttr("a", "href")
		allSevenSeasReleases = append(allSevenSeasReleases, temp)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://sevenseasentertainment.com/release-dates/")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allSevenSeasReleases)
}

func collectDarkHorseReleases() {
	var allDarkHorseReleases []Manga
	year, month, _ := time.Now().Date()

	collector := colly.NewCollector(
		colly.AllowedDomains("www.darkhorse.com", "darkhorse.com"),
	)

	collector.OnHTML(".list_item", func(element *colly.HTMLElement) {
		temp := Manga{}
		temp.Name = element.ChildText("a.product_link")
		temp.Image = element.ChildAttr("a.product_link img", "src")
		temp.Link = element.ChildAttr("a.product_link", "href")
		allDarkHorseReleases = append(allDarkHorseReleases, temp)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	darkHorseUrl := fmt.Sprintf("https://www.darkhorse.com/Books/Browse/Manga---%s+%d-%s+%d/P9wdwkt8", month, int(year), month, int(year))
	collector.Visit(darkHorseUrl)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allDarkHorseReleases)

}

func collectKodanshaReleases() {
	var allKodanshaReleases []Manga
	_, month, _ := time.Now().Date()

	collector := colly.NewCollector(
		colly.AllowedDomains("www.kodansha.us", "kodansha.us"),
	)

	collector.OnHTML(".calendar__day", func(element *colly.HTMLElement) {
		releaseDate := element.ChildText("h3.title.title--discovery")
		releaseDate = strings.Split(releaseDate, "/")[0]

		monthOfRelease, err := strconv.Atoi(releaseDate)
		if err != nil {
			fmt.Print("could not parse date")
			return
		}

		if monthOfRelease != int(month) {
			return
		}

		element.ForEach("div.card.book-card-small", func(_ int, el *colly.HTMLElement) {
			temp := Manga{}
			temp.Name = el.ChildText(".card__link")
			temp.Image = el.ChildAttr(".l-frame.product-image img", "src")
			temp.Link = el.ChildAttr("a.card__link", "href")
			allKodanshaReleases = append(allKodanshaReleases, temp)
		})
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	kodUrl := fmt.Sprintf("https://kodansha.us/manga/calendar")
	collector.Visit(kodUrl)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allKodanshaReleases)
}

func collectVizReleases() {
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

	pwd, _ := os.Getwd()
	s := fmt.Sprintf("pages/viz-%d-%d-%d", int(year), int(month), int(day))
	fmt.Println(s)
	collector.Visit("file://" + path.Join(pwd, s))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allVizReleases)
}

func main() {
	collectVizReleases()
	//collectYenPressReleases()
	//collectSevenSeasReleases()
	//collectDarkHorseReleases()
	//collectKodanshaReleases()
}
