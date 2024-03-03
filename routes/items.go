package routes

import (
	"html/template"
	"net/http"
	"strconv"

	db "github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

func Items(_html *template.Template) *Router {
	html = _html

	return NewRouter(
		Item,
		"/items",
		RouterActions{
			GetAll:  GetAllItems,
			GetByID: nil,
			Post:    nil,
			Put:     nil,
			Delete:  nil,
		},
	)
}

func GetAllItems(_ *http.Request) *web.Response {
	result, err := db.GetAllItems()
	if err != nil {
		panic(err)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"entity-list.html",
		result,
		nil,
	)
}

func GetItemByID(r *http.Request) *web.Response {
	var id int
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	result, err := db.GetItemByID(id)

	if err != nil {
		return web.Error(http.StatusInternalServerError, err, nil)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"item-by-id.html",
		result,
		nil,
	)
}
