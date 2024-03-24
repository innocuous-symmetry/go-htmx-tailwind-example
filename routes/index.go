package routes

import (
	"html/template"
	"net/http"
	// "github.com/innocuous-symmetry/moving-mgmt/"
	db "github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

var html *template.Template

func HomePage(r *http.Request) *web.Response {
	result, err := db.GetAllItems()
	if err != nil {
		panic(err)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"index.html",
		result,
		nil,
	)
}
