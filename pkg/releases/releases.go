package releases

import (
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
	Name  string `json:"name"`
	Image string `json:"image"`
	Link  string `json:"link"`
	Liked bool   `json:"liked"`
}

var location, _ = time.LoadLocation("UTC")
var year, month, day = time.Now().In(location).Date()

func CollectYenPressReleases() []Manga {
	var allYenReleases []Manga

	collector := colly.NewCollector()

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	collector.WithTransport(t)

	collector.OnHTML(".book-shelf-title-grid", func(element *colly.HTMLElement) {
		temp := Manga{}
		temp.Name = element.ChildText(".book-detail h2")
		temp.Image = element.ChildAttr("img", "src")
		temp.Link = "https://yenpress.com" + element.ChildAttr(".book-detail-links a", "href")
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

func CollectSevenSeasReleases() []Manga {
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

func CollectDarkHorseReleases() []Manga {
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

func CollectKodanshaReleases() []Manga {
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
		temp.Liked = false
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

func CollectTokyoPopReleases() []Manga {
	var allTokyoPopReleases []Manga
	_, month, _ := time.Now().Date()

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	collector := colly.NewCollector()
	collector.WithTransport(t)

	collector.OnHTML(".release-month", func(element *colly.HTMLElement) {
		fmt.Println("hallo")
		monthOfRelease := element.ChildText(".release-month-label")
		fmt.Println(monthOfRelease, month.String())
		if monthOfRelease != month.String() {
			return
		}

		temp := Manga{}
		temp.Name = element.ChildText(".rs-item-title")
		temp.Image = element.ChildAttr(".rs-item-image", "data-src")
		temp.Link = "bed"
		allTokyoPopReleases = append(allTokyoPopReleases, temp)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	pwd, _ := os.Getwd()
	s := fmt.Sprintf("pages/tokyopop-%d-%d-%d.html", int(year), int(month), int(day))
	collector.Visit("file://" + path.Join(pwd, s))

	return allTokyoPopReleases
}
