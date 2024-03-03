package routes

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

type Entity int

const (
	Item Entity = iota
	Box
	BoxItem
)

type RouterActions struct {
	GetAll 	func(r *http.Request) *web.Response
	GetByID func(r *http.Request) *web.Response
	Post 	func(r *http.Request) *web.Response
	Put 	func(r *http.Request) *web.Response
	Delete 	func(r *http.Request) *web.Response
}

type Router struct {
	Entity Entity
	Path string
	GetAll 		func(r *http.Request) *web.Response
	GetByID 	func(r *http.Request) *web.Response
	Post 		func(r *http.Request) *web.Response
	Put 		func(r *http.Request) *web.Response
	Delete 		func(r *http.Request) *web.Response
}

func NewRouter(entity Entity, path string, actions RouterActions) *Router {
	return &Router{
		Entity: entity,
		Path: path,
		GetAll: actions.GetAll,
		GetByID: actions.GetByID,
		Post: actions.Post,
		Put: actions.Put,
		Delete: actions.Delete,
	}
}
