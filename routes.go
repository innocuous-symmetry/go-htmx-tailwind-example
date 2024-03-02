package main

import (
	"net/http"
	db "github.com/innocuous-symmetry/moving-mgmt/db"

	"github.com/jritsema/gotoolbox/web"
)

func index(r *http.Request) *web.Response {
	result, err := db.GetAllItems()
	if err != nil {
		panic(err)
	}

	return web.HTML(http.StatusOK, html, "index.html", result, nil)
}
