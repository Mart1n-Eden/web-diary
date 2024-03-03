package handler

import (
	"net/http"
	"github.com/Mart1n-Eden/web-diary/internal/template"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if !limmiter.Allow() {
		http.Error(w, "Sorry((", http.StatusTooManyRequests)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "qwerty" {
			template.Tpl_admin.Execute(w, nil)
		} else {
			data := struct {
				ErrorMessage string
			}{
				ErrorMessage: "Неверный логин или пароль",
			}
			err := template.Tpl_login.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	} else {
		http.Error(w, "!!!!", http.StatusForbidden)
	}
}
