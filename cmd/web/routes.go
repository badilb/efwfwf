package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/news", app.showNews)
	mux.HandleFunc("/news/delete", app.deleteNews)
	mux.HandleFunc("/news/create", app.createNews)
	mux.HandleFunc("/news/creationPage", app.creationPage)
	mux.HandleFunc("/news/sport", app.filterCategory("Sports"))
	mux.HandleFunc("/news/politics", app.filterCategory("Politics"))
	mux.HandleFunc("/news/world", app.filterCategory("World"))
	mux.HandleFunc("/news/science", app.filterCategory("Science"))
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
