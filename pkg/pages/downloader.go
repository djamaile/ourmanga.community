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

package pages

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/robfig/cron"
)

type Site struct {
	Url  string
	Name string
}

type Date struct {
	location *time.Location
	year     int
	month    time.Month
	day      int
}

type Downloader struct {
	sites []Site
	date  Date
}

func (d *Downloader) removePages() {
	os.RemoveAll("pages/*")
}

func (d *Downloader) fetchPage(s Site) ([]byte, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(s.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return html, nil
}

func (d *Downloader) writePage(html []byte, s Site) error {
	fileName := fmt.Sprintf("pages/%v-%v-%v-%v.html", s.Name, d.date.year, int(d.date.month), d.date.day)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(html)
	if err != nil {
		return err
	}
	return nil
}

func (d *Downloader) refreshPages() error {
	d.date = newDate()
	d.sites = getSites(d.date)
	d.removePages()
	for _, s := range d.sites {
		html, err := d.fetchPage(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = d.writePage(html, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func newDate() Date {
	location, _ := time.LoadLocation("UTC")
	year, month, day := time.Now().In(location).Date()

	return Date{
		location: location,
		year:     year,
		month:    month,
		day:      day,
	}
}

func getSites(date Date) []Site {
	viz := Site{Name: "viz", Url: fmt.Sprintf("https://www.viz.com/calendar/%v/%v", date.year, int(date.month))}
	yenpress := Site{Name: "yenpress", Url: "https://yenpress.com/new-releases/"}
	sevenseas := Site{Name: "sevenseas", Url: "https://sevenseasentertainment.com/release-dates/"}
	darkhorse := Site{Name: "darkhorse", Url: fmt.Sprintf("https://www.darkhorse.com/Books/Browse/Manga---%v+%v-%v+%v/P9wdwkt8", date.month, date.year, date.month, date.year)}
	kodansha := Site{Name: "kodansha", Url: "https://kodansha.us/manga/calendar"}
	tokyopop := Site{Name: "tokyopop", Url: "https://www.tokyopop.com/upcoming"}
	square := Site{Name: "square", Url: "https://squareenixmangaandbooks.square-enix-games.com/en-us/release-calendar"}

	return []Site{viz, yenpress, sevenseas, darkhorse, kodansha, tokyopop, square}
}

func StartPagesJob() {
	downloader := &Downloader{}

	// initial downloading action on startup
	err := downloader.refreshPages()
	if err != nil {
		fmt.Printf("Error trying to download pages: %v", err)
	}

	c := cron.New()
	c.AddFunc("@hourly", func() {
		fmt.Println("Running cronjob to fetch pages")
		err := downloader.refreshPages()
		if err != nil {
			c.Stop()
			fmt.Printf("Error trying to download pages: %v", err)
		}
	})
	c.Start()
}
