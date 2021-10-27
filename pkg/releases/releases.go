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
	"fmt"
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
	releases, err := NewReleaseFetcher(".manga-books article", url, func(element *colly.HTMLElement) Manga {
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

func CollectTokyoPopReleases() []Manga {
	url := toLocalPagesPath("tokyopop")
	releases, err := NewReleaseFetcher(".release-month", url, func(element *colly.HTMLElement) Manga {
		monthOfRelease := element.ChildText(".release-month-label")
		fmt.Println(monthOfRelease, month.String())
		if monthOfRelease != month.String() {
			return Manga{}
		}

		temp := Manga{}
		temp.Name = element.ChildText(".rs-item-title")
		temp.Image = element.ChildAttr(".rs-item-image", "data-src")
		temp.Link = "bed"
		return temp
	}).Fetch()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return releases
}

func toLocalPagesPath(name string) string {
	return toPagesPath("pages", name)
}

func toPagesPath(rootDir string, name string) string {
	return fmt.Sprintf("%s/%s-%d-%d-%d.html", rootDir, name, int(year), int(month), int(day))
}
