package main

import (
	"html/template"
	"news-project/pkg/models"
	"path/filepath"
)

type templateData struct {
	News *models.News

	Flash          string
	NewsArray      []*models.News
	Category       string
	SuccessMessage string // Добавлено поле для сообщения об успехе
	ErrorMessage   string // Добавлено поле для сообщения об ошибке
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
