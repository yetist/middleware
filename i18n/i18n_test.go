// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package i18n

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Version(t *testing.T) {
	Convey("Check package version", t, func() {
		So(Version(), ShouldEqual, _VERSION)
	})
}

func Test_I18n(t *testing.T) {
	Convey("Use i18n middleware", t, func() {
		Convey("No langauge", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			m := martini.New()
			m.Use(I18n())
		})

		Convey("Languages and names not match", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			m := martini.New()
			m.Use(I18n(Options{
				Domain: "messages",
			}))
		})

		Convey("Invalid directory", func() {
			defer func() {
				So(recover(), ShouldBeNil)
			}()

			m := martini.New()
			m.Use(I18n(Options{
				Directory: "404",
			}))
		})

		Convey("With correct options", func() {
			m := martini.Classic()
			m.Use(I18n())

			m.Get("/foobar", func() {
			})

			resp := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/foobar", nil)
			So(err, ShouldBeNil)
			m.ServeHTTP(resp, req)
		})

		Convey("Set by redirect of URL parameter", func() {
			m := martini.Classic()
			m.Use(I18n(Options{Parameter: "lang"}))
			m.Get("/foobar", func() {
			})

			resp := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/foobar?lang=en-us", nil)
			So(err, ShouldBeNil)
			req.RequestURI = "/foobar?lang=en-us"
			m.ServeHTTP(resp, req)
		})

		Convey("Set by Accept-Language", func() {
			m := martini.Classic()
			m.Use(I18n(Options{}))
			m.Get("/foobar", func(l Locale) {
				So(l.Lang, ShouldEqual, "en-us")
			})

			resp := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/foobar", nil)
			So(err, ShouldBeNil)
			req.Header.Set("Accept-Language", "en-US")
			m.ServeHTTP(resp, req)
		})
	})
}
