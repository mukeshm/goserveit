package main

import "net/http"

type Route struct {
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
	Name        string
}

type Routes []Route

var routes = Routes{
	Route{
		"/",
		http.MethodGet,
		indexHandler,
		"Index",
	},
}
