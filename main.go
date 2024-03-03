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

	//go:embed css/output.css
	css embed.FS

	//parsed templates
	html *template.Template
)

func main() {

	//exit process immediately upon sigterm
	handleSigTerms()
	i, err := db.SeedDB()
	if err != nil {
		panic(err)
	}

	fmt.Printf("seeded db with %d records\n", i)

	//parse templates
	html, err = web.TemplateParseFSRecursive(templateFS, ".html", true, nil)
	if err != nil {
		panic(err)
	}

	//add routes
	router := http.NewServeMux()
	// router.Handle("/css/output.css", http.FileServer(http.FS(css)))

	// router.Handle("/company/add", web.Action(companyAdd))
	// router.Handle("/company/add/", web.Action(companyAdd))

	router.Handle("/", web.Action(routes.HomePage))
	router.Handle("/items", web.Action(routes.Items(html).GetAll))
	router.Handle("/items/{id}", web.Action(routes.Items(html).GetByID))
	router.Handle("/boxes", web.Action(routes.Boxes(html).GetAll))
	router.Handle("/unrelated", web.Action(func(r *http.Request) *web.Response {
		return web.HTML(http.StatusOK, html, "row-edit.html", nil, nil)
	}))

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
