package routes

import (
	"html/template"
	"net/http"

	"github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

type BoxActions struct {
	GetAll  func(r *http.Request) *web.Response
	GetByID func(r *http.Request) *web.Response
	Post    func(r *http.Request) *web.Response
	Put     func(r *http.Request) *web.Response
	Delete  func(r *http.Request) *web.Response
}

func Boxes(_html *template.Template) *BoxActions {
	html = _html

	return &BoxActions{
		GetAll:  GetAllBoxes,
		GetByID: nil,
		Post:    nil,
		Put:     nil,
		Delete:  nil,
	}
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
