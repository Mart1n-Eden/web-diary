package handler

import (
	"net/http"
	"github.com/Mart1n-Eden/web-diary/internal/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if !limmiter.Allow() {
		http.Error(w, "Sorry((", http.StatusTooManyRequests)
	}
	template.Tpl_login.Execute(w, nil)
}