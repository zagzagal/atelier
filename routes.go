package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		Index,
	},
	Route{
		"ApiIndex",
		"Get",
		"/api",
		ApiIndex,
	},
	Route{
		"ItemShow",
		"Get",
		"/api/item/{itemID}",
		ItemShow,
	},
	Route{
		"ItemList",
		"Get",
		"/api/item",
		ItemList,
	},
	Route{
		"PathShow",
		"Get",
		"/api/path/{start}/{dest}",
		PathShow,
	},
	Route{
		"ItemCreate",
		"POST",
		"/api/item",
		ItemCreate,
	},
}
