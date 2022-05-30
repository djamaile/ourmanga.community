// Copyright 2021 Djamaile Rahamat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package releases

import (
	"encoding/json"
	"fmt"
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
type TokyoPopManga struct {
	ReleaseDate string `json:"releaseDate"`
}

func CollectYenPressReleases() []Manga {
	url := toLocalPagesPath("yenpress")
	releases, err := NewReleaseFetcher(".book-shelf-title-grid", url, func(element *colly.HTMLElement) Manga {
		temp := Manga{}
		temp.Name = element.ChildText(".book-detail h2")
		temp.Image = element.ChildAttr("img", "src")
		temp.Link = "https://yenpress.com" + element.ChildAttr(".book-detail-links a", "href")
		return temp
	}).Fetch()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return releases
}

func CollectSevenSeasReleases() []Manga {
	url := toLocalPagesPath("sevenseas")
	releases, err := NewReleaseFetcher("div[style='float: left; margin: 0 3px 10px 6px; width: 134px; height: 189px; background: #CECECE;']", url, func(element *colly.HTMLElement) Manga {
		temp := Manga{}
		temp.Name = element.ChildAttr("a", "title")
		temp.Image = element.ChildAttr("img", "src")
		temp.Link = element.ChildAttr("a", "href")
		return temp
	}).Fetch()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return releases
}

func CollectDarkHorseReleases() []Manga {
	url := toLocalPagesPath("darkhorse")
	releases, err := NewReleaseFetcher(".list_item", url, func(element *colly.HTMLElement) Manga {
		temp := Manga{}
		temp.Name = element.ChildText("a.product_link")
		temp.Image = element.ChildAttr("a.product_link img", "src")
		temp.Link = element.ChildAttr("a.product_link", "href")
		return temp
	}).Fetch()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return releases
}

func CollectKodanshaReleases() []Manga {
	url := toLocalPagesPath("kodansha")
	releases, err := NewReleaseFetcher("div.card.book-card-small", url, func(element *colly.HTMLElement) Manga {
		temp := Manga{}
		temp.Name = element.ChildText(".card__link")
		temp.Image = element.ChildAttr(".l-frame.product-image img", "src")
		temp.Link = element.ChildAttr("a.card__link", "href")
		return temp
	}).Fetch()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return releases
}

func CollectVizReleases() []Manga {
	url := toLocalPagesPath("viz")
	releases, err := NewReleaseFetcher("#manga-books article", url, func(element *colly.HTMLElement) Manga {
		temp := Manga{}
		temp.Name = element.ChildText(".color-off-black")
		temp.Image = element.ChildAttr("a.product-thumb img", "data-original")
		temp.Link = "https://viz.com" + element.ChildAttr("a.product-thumb", "href")
		temp.Liked = false
		return temp
	}).Fetch()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return releases
}

// CollectTokyoPopReleases Tokyo pop release date is hiding in a json object within an attribute
// We need that date to check if it is only release for the wanted date
// We could give the Fetcher functionality a variable to know when it should stop
// TODO: clean this up
func CollectTokyoPopReleases() []Manga {
	neededManga := 0
	var location, _ = time.LoadLocation("UTC")
	var _, month, _ = time.Now().In(location).Date()
	url := toLocalPagesPath("tokyopop")
	releases, err := NewReleaseFetcher(".release-cal-item", url, func(element *colly.HTMLElement) Manga {
		var tokyoPopManga TokyoPopManga
		json.Unmarshal([]byte(element.ChildAttr(".rs-item-custom-fields", "data-custom-content")), &tokyoPopManga)
		releaseDate := fmt.Sprint(tokyoPopManga.ReleaseDate)
		convertedMonthToInt, err := strconv.Atoi(strings.Split(releaseDate, "/")[0])
		if err != nil {
			fmt.Println(err)
		}
		monthOfRelease := time.Month(convertedMonthToInt)
		if monthOfRelease == month {
			neededManga += 1
			manga := Manga{
				Name:  element.ChildAttr(".rs-item-thumbnail a", "data-title"),
				Image: element.ChildAttr(".rs-item-image-wrapper img", "data-src"),
				Link:  fmt.Sprintf("https://www.tokyopop.com%s", element.ChildAttr(".rs-item-details a", "href")),
				Liked: false,
			}
			return manga
		}
		return Manga{}
	}).Fetch()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return releases[:neededManga]
}

func toLocalPagesPath(name string) string {
	return toPagesPath("pages", name)
}

func toPagesPath(rootDir string, name string) string {
	var location, _ = time.LoadLocation("UTC")
	var year, month, day = time.Now().In(location).Date()
	return fmt.Sprintf("%s/%s-%d-%d-%d.html", rootDir, name, int(year), int(month), int(day))
}
