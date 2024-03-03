package routes

import (
	"net/http"

	"github.com/innocuous-symmetry/moving-mgmt/db"
	"github.com/jritsema/gotoolbox/web"
)

type RouterActions struct {
	GetAll 	func(r *http.Request) *web.Response
	GetByID func(r *http.Request) *web.Response
	Post 	func(r *http.Request) *web.Response
	Put 	func(r *http.Request) *web.Response
	Delete 	func(r *http.Request) *web.Response
}

type Router struct {
	Entity 		db.EntityLabel
	GetAll 		func(r *http.Request) *web.Response
	GetByID 	func(r *http.Request) *web.Response
	Post 		func(r *http.Request) *web.Response
	Put 		func(r *http.Request) *web.Response
	Delete 		func(r *http.Request) *web.Response
}

func NewRouter(entity db.EntityLabel, actions RouterActions) *Router {
	return &Router{
		Entity: entity,
		GetAll: actions.GetAll,
		GetByID: actions.GetByID,
		Post: actions.Post,
		Put: actions.Put,
		Delete: actions.Delete,
	}
}
