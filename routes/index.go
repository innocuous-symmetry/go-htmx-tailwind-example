package routes

import (
	"fmt"
	"html/template"
	"net/http"

	// "github.com/innocuous-symmetry/moving-mgmt/"
	db "github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

var html *template.Template

func HomePage(r *http.Request) *web.Response {
	result, err := db.GetAllItems()

	fmt.Println(result)

	if err != nil {
		return web.Error(http.StatusNotFound, err, nil)
	}

	items := []db.Item{}
	for result.Next() {
		item := db.Item{}
		err = db.ParseItem(&item, result.Scan)

		fmt.Println("name", item.Name)

		if err != nil {
			return web.Error(http.StatusInternalServerError, err, nil)
		}

		items = append(items, item)
	}

	return web.HTML(
		http.StatusOK,
		html,
		"index.html",
		items,
		nil,
	)
}
