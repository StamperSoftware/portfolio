﻿package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) Page(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}
