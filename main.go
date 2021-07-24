package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

type Manga struct {
	Name  string
	Image string
	Link  string
}

var year, month, day = time.Now().Date()

func collectYenPressReleases() []Manga {
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
	s := fmt.Sprintf("pages/yenpress-%d-%d-%d.html", int(year), int(month), int(day))
	collector.Visit("file://" + path.Join(pwd, s))

	return allYenReleases
}

func collectSevenSeasReleases() []Manga {
	var allSevenSeasReleases []Manga

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	collector := colly.NewCollector()
	collector.WithTransport(t)

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

	pwd, _ := os.Getwd()
	s := fmt.Sprintf("pages/sevenseas-%d-%d-%d.html", int(year), int(month), int(day))
	collector.Visit("file://" + path.Join(pwd, s))

	return allSevenSeasReleases
}

func collectDarkHorseReleases() []Manga {
	var allDarkHorseReleases []Manga
	year, month, _ := time.Now().Date()

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	collector := colly.NewCollector()
	collector.WithTransport(t)

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

	pwd, _ := os.Getwd()
	s := fmt.Sprintf("pages/darkhorse-%d-%d-%d.html", int(year), int(month), int(day))
	collector.Visit("file://" + path.Join(pwd, s))

	return allDarkHorseReleases
}

func collectKodanshaReleases() []Manga {
	var allKodanshaReleases []Manga
	_, month, _ := time.Now().Date()

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	collector := colly.NewCollector()
	collector.WithTransport(t)

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

	pwd, _ := os.Getwd()
	s := fmt.Sprintf("pages/kodansha-%d-%d-%d.html", int(year), int(month), int(day))
	collector.Visit("file://" + path.Join(pwd, s))

	return allKodanshaReleases
}

func collectVizReleases() []Manga {
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
	s := fmt.Sprintf("pages/viz-%d-%d-%d.html", int(year), int(month), int(day))
	collector.Visit("file://" + path.Join(pwd, s))

	return allVizReleases
}

func vizHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(collectVizReleases())
}

func yenHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(collectYenPressReleases())
}

func sevenSeasHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(collectSevenSeasReleases())
}

func darkHorseHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(collectDarkHorseReleases())
}

func kodanshaHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(collectKodanshaReleases())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/releases/viz", vizHandler)
	r.HandleFunc("/releases/yenpress", yenHandler)
	r.HandleFunc("/releases/sevenseas", sevenSeasHandler)
	r.HandleFunc("/releases/darkhorse", darkHorseHandler)
	r.HandleFunc("/releases/kodansha", kodanshaHandler)

	log.Println("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
