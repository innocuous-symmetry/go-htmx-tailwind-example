package routes

import (
	"html/template"
	"net/http"
	"strconv"

	db "github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

type ItemActions struct {
	Get    func(r *http.Request) *web.Response
	GetAll func(r *http.Request) *web.Response
	Edit   func(r *http.Request) *web.Response
	Delete func(r *http.Request) *web.Response
	Save   func(r *http.Request) *web.Response
}

func Items(_html *template.Template) *ItemActions {
	html = _html

	return &ItemActions{
		Get:    Get,
		GetAll: GetAllItems,
		Edit:   EditItem,
		Delete: nil,
		Save:   Put,
	}
}

func Get(r *http.Request) *web.Response {
	_, count := web.PathLast(r)

	if count == 0 {
		return GetAllItems(r)
	} else {
		return GetItemByID(r)
	}
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

func EditItem(r *http.Request) *web.Response {
	idFromPath, _ := web.PathLast(r)

	id, err := strconv.ParseInt(idFromPath, 10, 64)
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	result, err := db.GetItemByID(int(id))

	if err != nil {
		return web.Error(http.StatusInternalServerError, err, nil)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"entity-edit.html",
		result,
		nil,
	)
}

func GetItemByID(r *http.Request) *web.Response {
	idFromPath, _ := web.PathLast(r)

	id, err := strconv.ParseInt(idFromPath, 10, 64)
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	result, err := db.GetItemByID(int(id))

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

func Put(r *http.Request) *web.Response {
	id, _ := web.PathLast(r)

	err := r.ParseForm()
	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	name := r.Form.Get("name")
	stage := r.Form.Get("stage")
	category := r.Form.Get("category")
	description := r.Form.Get("description")
	notes := r.Form.Get("notes")
	// id := r.Form.Get("id")

	item := db.Item{
		ID: func() int {
			idInt, err := strconv.ParseInt(id, 10, 64)

			if err != nil {
				return -1
			}

			return int(idInt)
		}(),

		Name:        name,
		Description: &description,
		Notes:       &notes,

		Stage: func() db.PackingStage {
			stageInt, _ := strconv.Atoi(stage)
			return db.PackingStage(stageInt)
		}(),

		Category: func() db.Category {
			categoryInt, _ := strconv.Atoi(category)
			return db.Category(categoryInt)
		}(),
	}

	_, err = db.PutItem(item)

	if err != nil {
		return web.Error(http.StatusInternalServerError, err, nil)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"entity-row.html",
		item,
		nil,
	)
}

func Delete(r *http.Request) *web.Response {
	idFromPath, _ := web.PathLast(r)
	id, err := strconv.ParseInt(idFromPath, 10, 64)

	if err != nil {
		return web.Error(http.StatusBadRequest, err, nil)
	}

	_, err = db.DeleteItem(int(id))

	if err != nil {
		return web.Error(http.StatusInternalServerError, err, nil)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"entity-row.html",
		nil,
		nil,
	)
}
