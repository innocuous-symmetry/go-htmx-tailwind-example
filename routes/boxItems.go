package routes

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

type BoxItemActions struct {
	Get        func(r *http.Request) *web.Response
	GetAll     func(r *http.Request) *web.Response
	GetByID    func(r *http.Request) *web.Response
	GetByBoxID func(r *http.Request) *web.Response
}

func BoxItems(_html *template.Template) *BoxItemActions {
	html = _html

	return &BoxItemActions{
		Get:     BoxItemsHandler,
		GetAll:  nil,
		GetByID: nil,
	}
}

func BoxItemsHandler(r *http.Request) *web.Response {
	switch r.Method {

	case http.MethodGet:
		if r.URL.Query().Has("boxid") {
			return GetBoxItemsByBoxID(r)
		}

		_, count := web.PathLast(r)
		if count == 1 {
			return GetAllBoxItems(r)
		} else {
			return GetBoxItemByID(r)
		}
	default:
		return nil

	}
}

func GetAllBoxItems(_ *http.Request) *web.Response {
	items, err := db.GetAllBoxItems()
	if err != nil {
		return web.Error(
			http.StatusBadRequest,
			err,
			nil,
		)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"box-items/box-item-list.html",
		items,
		nil,
	)
}

func GetBoxItemsByBoxID(r *http.Request) *web.Response {
	boxID := r.URL.Query().Get("boxid")

	if id, err := strconv.ParseInt(boxID, 10, 64); err != nil {
		return web.Error(
			http.StatusBadRequest,
			err,
			nil,
		)
	} else {
		items, err := db.GetBoxItemsByBoxID(int(id))
		if err != nil {
			return web.Error(
				http.StatusNotFound,
				err,
				nil,
			)
		}

		return web.HTML(
			http.StatusOK,
			html,
			"box-items/box-item-list.html",
			items,
			nil,
		)

	}
}

func GetBoxItemByID(_ *http.Request) *web.Response {
	return nil
}
