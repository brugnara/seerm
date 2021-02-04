package main

import (
	"encoding/json"
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
	Note  string
}

func init() {
}

func (c customer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})`)
	if m := re.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		// handle
		var c customer
		res := db.First(&c, "uuid = ?", m[1])

		if res.RowsAffected == 0 {
			tpls.ExecuteTemplate(w, "customer.gohtml", nil)
			return
		}

		switch r.Method {
		case http.MethodGet:
			tpls.ExecuteTemplate(w, "customer.gohtml", &c)
			return
		case http.MethodDelete:
			db.Delete(&c)
			return
		case http.MethodPatch:
			var dat map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&dat); err != nil {
				panic(err)
			}
			db.Model(&c).Update(
				"note", dat["note"].(string),
			)
			return
		case http.MethodPost:
			db.Model(&c).Update(
				"name", r.FormValue("name"),
			).Update(
				"email", r.FormValue("email"),
			)
		}

		http.Redirect(w, r, "/c/"+c.UUID, http.StatusSeeOther)
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
	// tpls.ExecuteTemplate(w, "customer.gohtml", &cst)
	http.Redirect(w, r, "/c/"+cst.UUID, http.StatusSeeOther)
}
