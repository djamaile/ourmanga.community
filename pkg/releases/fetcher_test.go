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
	"github.com/gocolly/colly"
	"net"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_releaseFetcher_Fetch(t *testing.T) {
	transport := &http.Transport{
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
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	type fields struct {
		http          *http.Transport
		url           string
		querySelector string
		mapper        elementMapper
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Manga
		wantErr bool
	}{
		{
			name: "should fetch a single manga from test page",
			fields: fields{
				http:          transport,
				url:           "testdata/test.html",
				querySelector: ".test_div",
				mapper: func(element *colly.HTMLElement) Manga {
					temp := Manga{}
					temp.Name = element.ChildText(".title")
					temp.Image = element.ChildAttr(".item img", "src")
					temp.Link = element.ChildAttr("a.link", "href")
					temp.Liked = false
					return temp
				},
			},
			want: []Manga{
				{
					Name:  "Test title",
					Image: "test.png",
					Link:  "https://test.nl",
					Liked: false,
				},
			},
		},
		{
			name: "should return an error when fetching manga from a non-existing path",
			fields: fields{
				http: transport,
				url:  "testdata/not_existing_page.html",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := releaseFetcher{
				http:          tt.fields.http,
				url:           tt.fields.url,
				querySelector: tt.fields.querySelector,
				mapper:        tt.fields.mapper,
			}
			got, err := r.Fetch()
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
