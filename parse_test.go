// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPI = []route{
	// Objects
	{"POST", "/1/classes/:className"},
	{"GET", "/1/classes/:className/:objectId"},
	{"PUT", "/1/classes/:className/:objectId"},
	{"GET", "/1/classes/:className"},
	{"DELETE", "/1/classes/:className/:objectId"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/:objectId"},
	{"PUT", "/1/users/:objectId"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/:objectId"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/:objectId"},
	{"PUT", "/1/roles/:objectId"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/:objectId"},

	// Files
	{"POST", "/1/files/:fileName"},

	// Analytics
	{"POST", "/1/events/:eventName"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/:objectId"},
	{"PUT", "/1/installations/:objectId"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/:objectId"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var (
	parseBeego       http.Handler
	parseBunRouter   http.Handler
	parseChi         http.Handler
	parseEcho        http.Handler
	parseFiber       *fiber.App
	parseGin         http.Handler
	parseGojiv2      http.Handler
	parseGoRestful   http.Handler
	parseGorillaMux  http.Handler
	parseHttpRouter  http.Handler
	parseHttpTreeMux http.Handler
	parseMacaron     http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPI))

	calcMem("Beego", func() {
		parseBeego = loadBeego(parseAPI)
	})
	calcMem("BunRouter", func() {
		parseBunRouter = loadBunRouter(parseAPI)
	})
	calcMem("Chi", func() {
		parseChi = loadChi(parseAPI)
	})
	calcMem("Echo", func() {
		parseEcho = loadEcho(parseAPI)
	})
	calcMem("Fiber", func() {
		parseFiber = loadFiber(parseAPI)
	})
	calcMem("Gin", func() {
		parseGin = loadGin(parseAPI)
	})
	calcMem("Gojiv2", func() {
		parseGojiv2 = loadGojiv2(parseAPI)
	})
	calcMem("GoRestful", func() {
		parseGoRestful = loadGoRestful(parseAPI)
	})
	calcMem("GorillaMux", func() {
		parseGorillaMux = loadGorillaMux(parseAPI)
	})
	calcMem("HttpRouter", func() {
		parseHttpRouter = loadHttpRouter(parseAPI)
	})
	calcMem("HttpTreeMux", func() {
		parseHttpTreeMux = loadHttpTreeMux(parseAPI)
	})
	calcMem("Macaron", func() {
		parseMacaron = loadMacaron(parseAPI)
	})

	println()
}

// Static
func BenchmarkBeego_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBunRouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBunRouter, req)
}
func BenchmarkChi_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseChi, req)
}
func BenchmarkEcho_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkFiber_ParseStatic(b *testing.B) {
	benchFiberRequest(b, parseFiber, "GET", "/1/users")
}
func BenchmarkGin_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGojiv2_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGojiv2, req)
}
func BenchmarkGoRestful_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGoRestful, req)
}
func BenchmarkGorillaMux_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkMacaron_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseMacaron, req)
}

// One Param
func BenchmarkBeego_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBunRouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBunRouter, req)
}
func BenchmarkChi_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseChi, req)
}
func BenchmarkEcho_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkFiber_ParseParam(b *testing.B) {
	benchFiberRequest(b, parseFiber, "GET", "/1/classes/go")
}
func BenchmarkGin_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGojiv2_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGojiv2, req)
}
func BenchmarkGoRestful_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGoRestful, req)
}
func BenchmarkGorillaMux_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkMacaron_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseMacaron, req)
}

// Two Params
func BenchmarkBeego_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBunRouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBunRouter, req)
}
func BenchmarkChi_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseChi, req)
}
func BenchmarkEcho_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkFiber_Parse2Params(b *testing.B) {
	benchFiberRequest(b, parseFiber, "GET", "/1/classes/go/123456789")
}
func BenchmarkGin_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGojiv2_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGojiv2, req)
}
func BenchmarkGoRestful_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGoRestful, req)
}
func BenchmarkGorillaMux_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkHttpRouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkMacaron_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseMacaron, req)
}

// All Routes
func BenchmarkBeego_ParseAll(b *testing.B) {
	benchRoutes(b, parseBeego, parseAPI)
}
func BenchmarkBunRouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseBunRouter, parseAPI)
}
func BenchmarkChi_ParseAll(b *testing.B) {
	benchRoutes(b, parseChi, parseAPI)
}
func BenchmarkEcho_ParseAll(b *testing.B) {
	benchRoutes(b, parseEcho, parseAPI)
}
func BenchmarkFiber_ParseAll(b *testing.B) {
	benchFiberRoutes(b, parseFiber, parseAPI)
}
func BenchmarkGin_ParseAll(b *testing.B) {
	benchRoutes(b, parseGin, parseAPI)
}
func BenchmarkGojiv2_ParseAll(b *testing.B) {
	benchRoutes(b, parseGojiv2, parseAPI)
}
func BenchmarkGoRestful_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoRestful, parseAPI)
}
func BenchmarkGorillaMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseGorillaMux, parseAPI)
}
func BenchmarkHttpRouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseHttpRouter, parseAPI)
}
func BenchmarkHttpTreeMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseHttpTreeMux, parseAPI)
}
func BenchmarkMacaron_ParseAll(b *testing.B) {
	benchRoutes(b, parseMacaron, parseAPI)
}
