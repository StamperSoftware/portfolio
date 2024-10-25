package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

type application struct {
	templateMap map[string]*template.Template
}

func main() {

	app := application{
		templateMap: make(map[string]*template.Template),
	}

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting on port", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
