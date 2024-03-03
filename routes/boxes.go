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
		"boxes",
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
	result, err := db.GetAll("boxes")
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	boxes := []db.Box{}

	for result.Next() {
		box := db.Box{}
		err = db.ParseBox(&box, result.Scan)
		if err != nil {
			return web.Error(http.StatusInternalServerError, err, nil)
		}

		boxes = append(boxes, box)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"entity-list.html",
		result,
		nil,
	)
}
