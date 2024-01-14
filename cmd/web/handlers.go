package main

import (
	"NewsAituuu-main/pkg/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	s, err := app.news.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{NewsArray: s}
	files := []string{"./ui/html/home.page.tmpl", "./ui/html/base.layout.tmpl", "./ui/html/footer.partial.tmpl"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showNews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.news.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrorMessage) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := &templateData{News: s}
	files := []string{"./ui/html/show.page.tmpl", "./ui/html/base.layout.tmpl", "./ui/html/footer.partial.tmpl"}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) creationPage(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "create.page.tmpl", &templateData{})
}

func (app *application) updateNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")

	err = app.news.UpdateNewsById(id, title, content, category)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/news?id=%d", id), http.StatusSeeOther)
}

func (app *application) deleteNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	paramId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.news.DeleteNewsById(id)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) createNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")

	id, err := app.news.Insert(title, content, category)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/news?id=%d", id), http.StatusSeeOther)
}

func (app *application) filterCategory(category string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := app.renderCategoryPage(w, r, category)
		if err != nil {
			app.serverError(w, err)
			return
		}
	}
}

func (app *application) renderCategoryPage(w http.ResponseWriter, r *http.Request, category string) error {
	newsArray, err := app.news.GetByCategory(category)
	if err != nil {
		return err
	}
	app.render(w, r, "category.page.tmpl", &templateData{
		Category:  category,
		NewsArray: newsArray,
	})
	return nil
}
