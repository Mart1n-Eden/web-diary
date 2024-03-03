package handler

import (
	"net/http"
	"github.com/Mart1n-Eden/web-diary/internal/template"
	"github.com/Mart1n-Eden/web-diary/internal/model"
	"github.com/Mart1n-Eden/web-diary/internal/database"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	if !limmiter.Allow() {
		http.Error(w, "Sorry((", http.StatusTooManyRequests)
	}

	if r.Method == http.MethodPost {
		a := model.Article{Title: r.FormValue("title"), Body: []byte(r.FormValue("content"))}
		database.SaveArticle(a)
		template.Tpl_admin.Execute(w, nil)
	}
}