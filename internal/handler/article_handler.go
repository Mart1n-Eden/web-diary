package handler

import (
	"strconv"
	"net/http"
	"github.com/Mart1n-Eden/web-diary/internal/model"
	"github.com/Mart1n-Eden/web-diary/internal/database"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	if !limmiter.Allow() {
		http.Error(w, "Sorry((", http.StatusTooManyRequests)
	}

	param := r.URL.Query().Get("id")
	i, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}

	a := model.Article{}
	database.TakeArticle(&a,i)

	// template.Tpl_article.Execute(w, data)
	w.Write(a.Gluing())
}