package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	db "github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/innocuous-symmetry/moving-mgmt/util"
	"github.com/jritsema/gotoolbox/web"
)

func Items(_html *template.Template) *Router {
	html = _html

	return NewRouter(
		"items",
		RouterActions{
			GetAll:  GetAllItems,
			GetByID: GetItemByID,
			Post:    nil,
			Put:     nil,
			Delete:  nil,
		},
	)
}

func GetAllItems(_ *http.Request) *web.Response {
	result, err := db.GetAll("items")
	if err != nil {
		return web.Error(http.StatusNotFound, err, nil)
	}

	items := []db.Item{}

	for result.Next() {
		item := db.Item{}
		err = db.ParseItem(&item, result.Scan)
		if err != nil {
			fmt.Println(err.Error())
			return web.Error(http.StatusInternalServerError, err, nil)
		}

		items = append(items, item)
	}

	fmt.Println("items", items)

	return web.HTML(
		http.StatusOK,
		html,
		"entity-list.html",
		result,
		nil,
	)
}

func GetItemByID(r *http.Request) *web.Response {
	id, err := util.GetIDFromPath(r)
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	editMode, err := strconv.ParseBool(r.URL.Query().Get("edit"))
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	res, err := db.GetByID("items", id)

	item := db.Item{}
	err = db.ParseItem(&item, res.Scan)
	if err != nil {
		return web.Error(http.StatusInternalServerError, err, nil)
	}

	var tmpl string

	if editMode {
		tmpl = "entity-edit.html"
	} else {
		tmpl = "entity-row.html"
	}

	return web.HTML(
		http.StatusOK,
		html,
		tmpl,
		item,
		nil,
	)
}

func PutItem(r *http.Request) *web.Response {
	body := r.Body
	defer body.Close()

	bodyBytes := make([]byte, r.ContentLength)
	_, err := body.Read(bodyBytes)
	if err != nil {
		return web.Error(http.StatusInternalServerError, err, nil)
	}

	item, err := db.ParseEntityFromBytes(bodyBytes)
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	result, err := db.Put("items", item)
	if err != nil {
		return web.Error(http.StatusInternalServerError, err, nil)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"entity-row.html",
		result,
		nil,
	)
}
