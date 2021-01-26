package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

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
		tpls.ExecuteTemplate(w, "customer.gohtml", nil)
		return
	}
	// new
	if r.Method == http.MethodPost {
		r.ParseForm()
		cst := customer{
			UUID:  fmt.Sprintf("%s", uuid.NewV4()),
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
		}
		log.Println(cst)
		db.Create(&cst)
		tpls.ExecuteTemplate(w, "customer.gohtml", &cst)
		return
	}
	tpls.ExecuteTemplate(w, "customer-new.gohtml", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {

}
