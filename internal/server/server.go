package server

import (
	"net/http"
	"os"
	"os/signal"
	"github.com/Mart1n-Eden/web-diary/internal/database"
	"github.com/Mart1n-Eden/web-diary/internal/handler"
)

func Run() {
	database.Init()

	mux := http.NewServeMux()
	mux.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/admin", handler.AdminHandler)
	mux.HandleFunc("/save", handler.SaveHandler)
	mux.HandleFunc("/article", handler.ArticleHandler)

	http.ListenAndServe(":8888", mux)

	go func() {
		if err := http.ListenAndServe(":8888", mux); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	
	database.Close()
}
