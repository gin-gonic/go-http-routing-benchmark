package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// load functions of all routers
	routers = []struct {
		name string
		load func(routes []route) http.Handler
	}{
		{"Beego", loadBeego},
		{"BunRouter", loadBunRouter},
		{"Chi", loadChi},
		{"Echo", loadEcho},
		{"Gin", loadGin},
		{"Gojiv2", loadGojiv2},
		{"GoRestful", loadGoRestful},
		{"GorillaMux", loadGorillaMux},
		{"HttpRouter", loadHttpRouter},
		{"HttpTreeMux", loadHttpTreeMux},
		{"Macaron", loadMacaron},
	}

	// all APIs
	apis = []struct {
		name   string
		routes []route
	}{
		{"GitHub", githubAPI},
		{"GPlus", gplusAPI},
		{"Parse", parseAPI},
		{"Static", staticRoutes},
	}
)

func TestRouters(t *testing.T) {
	loadTestHandler = true

	for _, router := range routers {
		req, _ := http.NewRequest("GET", "/", nil)
		u := req.URL
		rq := u.RawQuery

		for _, api := range apis {
			r := router.load(api.routes)

			for _, route := range api.routes {
				w := httptest.NewRecorder()
				req.Method = route.method
				req.RequestURI = route.path
				u.Path = route.path
				u.RawQuery = rq
				r.ServeHTTP(w, req)
				if w.Code != 200 || w.Body.String() != route.path {
					t.Errorf(
						"%s in API %s: %d - %s; expected %s %s\n",
						router.name, api.name, w.Code, w.Body.String(), route.method, route.path,
					)
				}
			}
		}
	}

	loadTestHandler = false
}
