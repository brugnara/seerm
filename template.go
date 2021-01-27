package main

import (
	"fmt"
	"html/template"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpls *template.Template

type hero struct {
	Title    string
	Subtitle string
}

func init() {
	tpls = template.Must(
		template.New("").Funcs(
			template.FuncMap{
				"toHero": toHero,
				"concat": concat,
				"randID": randID,
			},
		).ParseGlob("./templates/*.gohtml"),
	)
}

func toHero(s string) (ret hero) {
	tmp := strings.Split(s, "|")
	if len(tmp) > 0 {
		ret.Title = tmp[0]
	}
	if len(tmp) > 1 {
		ret.Subtitle = tmp[1]
	}
	return
}

func concat(a, b string) string {
	return a + "|" + b
}

func randID() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}
