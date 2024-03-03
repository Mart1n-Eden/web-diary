package template

import (
	"html/template"
)

var (
	Tpl       = template.Must(template.ParseFiles("./view/index.html"))
	Tpl_login = template.Must(template.ParseFiles("./view/login.html"))
	Tpl_admin = template.Must(template.ParseFiles("./view/admin.html"))
)