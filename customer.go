package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type customer struct {
	gorm.Model
	UUID  string `gorm:"index"`
	Name  string
	Email string `gorm:"index"`
}

func init() {
}

func (c customer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})`)
	if m := re.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		// handle
		var c customer
		db.First(&c, "uuid = ?", m[1])

		if r.FormValue("delete") != "" {
			db.Delete(&c)
			tpls.ExecuteTemplate(w, "customer.gohtml", nil)
			return
		}

		tpls.ExecuteTemplate(w, "customer.gohtml", &c)
		return
	}
	// new
	if r.Method == http.MethodPost {
		c.post(w, r)
		return
	}
	if sx := strings.Split(r.URL.Path, "/"); sx[len(sx)-1] == "new" {
		tpls.ExecuteTemplate(w, "customer-new.gohtml", nil)
		return
	}
	// show all
	c.list(w, r)
}

func (c customer) list(w http.ResponseWriter, r *http.Request) {
	var cx []customer
	db.Limit(50).Find(&cx)
	tpls.ExecuteTemplate(w, "customers.gohtml", &cx)
}

func (c customer) post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cst := customer{
		UUID:  fmt.Sprintf("%s", uuid.NewV4()),
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}
	log.Println(cst)
	db.Create(&cst)
	tpls.ExecuteTemplate(w, "customer.gohtml", &cst)
}
