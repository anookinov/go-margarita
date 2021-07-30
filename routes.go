package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := Routes{
		Route{
			"getMembers",
			"GET",
			"/api/members",
			getMembers,
		},
		Route{
			"getMember",
			"GET",
			"/api/members/{id}",
			getMember,
		},
		Route{
			"createMember",
			"POST",
			"/api/members",
			createMember,
		},
		Route{
			"updateMember",
			"PUT",
			"/api/members/{id}",
			updateMember,
		},
		Route{
			"deleteMember",
			"DELETE",
			"/api/members/{id}",
			deleteMember,
		},
	}
	for _, r := range routes {
		router.Methods(r.Method).Path(r.Pattern).Name(r.Name).Handler(r.HandlerFunc)
	}
	return router
}