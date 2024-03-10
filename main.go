package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/innocuous-symmetry/moving-mgmt/routes"

	"github.com/jritsema/gotoolbox"
	"github.com/jritsema/gotoolbox/web"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS

	//parsed templates
	html *template.Template
)

func main() {

	//exit process immediately upon sigterm
	handleSigTerms()
	db.SeedDB()

	//parse templates
	var err error
	html, err = web.TemplateParseFSRecursive(templateFS, ".html", true, nil)
	if err != nil {
		panic(err)
	}

	//add routes
	router := http.NewServeMux()

	itemActions := routes.Items(html)
	boxActions := routes.Boxes(html)

	router.Handle("/items/edit", web.Action(itemActions.Edit))
	router.Handle("/items/delete", web.Action(itemActions.Delete))
	router.Handle("/items/save", web.Action(itemActions.Save))
	router.Handle("/items/edit/", web.Action(itemActions.Edit))
	router.Handle("/items/delete/", web.Action(itemActions.Delete))
	router.Handle("/items/delete/:id", web.Action(itemActions.Delete))
	router.Handle("/items/save/", web.Action(itemActions.Save))
	router.Handle("/items/save/:id", web.Action(itemActions.Save))

	router.Handle("/items/add", web.Action(itemActions.Post))
	router.Handle("/items/add/", web.Action(itemActions.Post))

	router.Handle("/items", web.Action(itemActions.Get))
	router.Handle("/boxes", web.Action(boxActions.GetAll))
	router.Handle("/items/", web.Action(itemActions.Get))
	router.Handle("/items/:id", web.Action(itemActions.Get))
	router.Handle("/boxes/", web.Action(boxActions.GetAll))
	router.Handle("/boxes/:id", web.Action(boxActions.GetAll))

	router.Handle("/", web.Action(routes.HomePage))
	router.Handle("/index.html", web.Action(routes.HomePage))

	//logging/tracing
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := tracing(nextRequestID)(logging(logger)(router))

	port := gotoolbox.GetEnvWithDefault("PORT", "8080")
	logger.Println("listening on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, middleware); err != nil {
		logger.Println("http.ListenAndServe():", err)
		os.Exit(1)
	}
}

func handleSigTerms() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
