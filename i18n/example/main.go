package main

import (
	"github.com/chai2010/gettext-go/gettext"
	"github.com/go-martini/martini"
	"github.com/yetist/middleware/i18n"
)

func __(msgid string) string {
	return gettext.PGettext("", msgid)
}

func main() {
	m := martini.Classic()
	m.Use(i18n.I18n(i18n.Options{
		Domain:    "example",
		Parameter: "lang",
	}))
	m.Get("/", func() string {
		return __("Hello world!")
	})
	m.Run()
}
