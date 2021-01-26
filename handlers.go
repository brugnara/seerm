package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var count int64
	var srcRes []customer
	var query string

	db.Model(&customer{}).Count(&count)

	query = r.FormValue("q")

	if q := query; q != "" {
		log.Println("Searching:", q)
		like := "%" + q + "%"
		db.Where(
			"name LIKE ? OR email LIKE ?", like, like,
		).Find(&srcRes)
	}

	if err := tpls.ExecuteTemplate(w, "index.gohtml", struct {
		Count      int64
		SearchRes  []customer
		SearchTerm string
	}{count, srcRes, query}); err != nil {
		log.Println(err)
	}
}
