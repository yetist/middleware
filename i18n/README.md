# i18n

i18n is a Martini middleware for internationalization. 

Support po/mo file, use [gettext-go](https://github.com/chai2010/gettext-go).

Usage
======

main.go

```
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
```

directories:

```
locale/
├── en_US
│   └── LC_MESSAGES
│       ├── example.mo
│       └── example.po
├── zh_CN
│   └── LC_MESSAGES
│       ├── example.mo
│       └── example.po
└── zh_TW
    └── LC_MESSAGES
        ├── example.mo
	└── example.po
```
