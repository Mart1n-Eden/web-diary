package handler

import (
	"strconv"
	"net/http"
	"golang.org/x/time/rate"

	"github.com/Mart1n-Eden/web-diary/internal/database"
	"github.com/Mart1n-Eden/web-diary/internal/template"
	"github.com/Mart1n-Eden/web-diary/internal/model"
)

var	limmiter = rate.NewLimiter(rate.Limit(100), 1)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if !limmiter.Allow() {
		http.Error(w, "Sorry((", http.StatusTooManyRequests)
	}

	var page int
	var err error

	pageStr := r.URL.Query().Get("id")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			panic(err)
		}
	}

	count := database.CountArticles()

	offset := page * 3
	pag := model.Pagination{}
	if page == 0 {
		pag.HasPrev = false
	} else {
		pag.HasPrev = true
	}
	if page >= count/3 {
		pag.HasNext = false
	} else {
		pag.HasNext = true
	}
	pag.Page = page
	pag.NextPage = page + 1
	pag.PrevPage = page - 1

	data := model.PageData{
		Articles:   database.GetTemplate(offset),
		Pagination: pag,
	}
	template.Tpl.Execute(w, data)
}