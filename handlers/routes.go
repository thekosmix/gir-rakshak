package handlers

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
		"GET",
		"/",
		Index,
	},
	Route{
		"AllUsers",
		"GET",
		"/api/admin/users",
		AllUsers,
	},
	Route{
		"ApproveUser",
		"POST",
		"/api/admin/users",
		ApproveUser,
	},
	Route{
		"UserLocation",
		"GET",
		"/api/admin/users/{userId}/locations",
		UserLocation,
	},
	Route{
		"RegisterUser",
		"POST",
		"/api/users/register",
		RegisterUser,
	},
	Route{
		"LoginUser",
		"POST",
		"/api/users/login",
		LoginUser,
	},
	Route{
		"UploadUserLocation",
		"POST",
		"/api/users/{userId}/locations",
		UploadUserLocation,
	},
}
