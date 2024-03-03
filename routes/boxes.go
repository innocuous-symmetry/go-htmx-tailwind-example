package routes

import (
	"html/template"
	"net/http"

	"github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

func Boxes(_html *template.Template) *Router {
	html = _html

	return NewRouter(
		Box,
		"/boxes",
		RouterActions{
			GetAll:  GetAllBoxes,
			GetByID: nil,
			Post:    nil,
			Put:     nil,
			Delete:  nil,
		},
	)
}

func GetAllBoxes(_ *http.Request) *web.Response {
	result, err := db.GetAllBoxes()
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"entity-list.html",
		result,
		nil,
	)
}
