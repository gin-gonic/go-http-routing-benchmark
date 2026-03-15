// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var gplusAPI = []route{
	// People
	{"GET", "/people/:userId"},
	{"GET", "/people"},
	{"GET", "/activities/:activityId/people/:collection"},
	{"GET", "/people/:userId/people/:collection"},
	{"GET", "/people/:userId/openIdConnect"},

	// Activities
	{"GET", "/people/:userId/activities/:collection"},
	{"GET", "/activities/:activityId"},
	{"GET", "/activities"},

	// Comments
	{"GET", "/activities/:activityId/comments"},
	{"GET", "/comments/:commentId"},

	// Moments
	{"POST", "/people/:userId/moments/:collection"},
	{"GET", "/people/:userId/moments/:collection"},
	{"DELETE", "/moments/:id"},
}

var (
	gplusBeego       http.Handler
	gplusBunRouter   http.Handler
	gplusChi         http.Handler
	gplusEcho        http.Handler
	gplusFiber       *fiber.App
	gplusGin         http.Handler
	gplusGojiv2      http.Handler
	gplusGoRestful   http.Handler
	gplusGorillaMux  http.Handler
	gplusHttpRouter  http.Handler
	gplusHttpTreeMux http.Handler
	gplusMacaron     http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	calcMem("Beego", func() {
		gplusBeego = loadBeego(gplusAPI)
	})
	calcMem("BunRouter", func() {
		gplusBunRouter = loadBunRouter(gplusAPI)
	})
	calcMem("Chi", func() {
		gplusChi = loadChi(gplusAPI)
	})
	calcMem("Echo", func() {
		gplusEcho = loadEcho(gplusAPI)
	})
	calcMem("Fiber", func() {
		gplusFiber = loadFiber(gplusAPI)
	})
	calcMem("Gin", func() {
		gplusGin = loadGin(gplusAPI)
	})
	calcMem("Gojiv2", func() {
		gplusGojiv2 = loadGojiv2(gplusAPI)
	})
	calcMem("GoRestful", func() {
		gplusGoRestful = loadGoRestful(gplusAPI)
	})
	calcMem("GorillaMux", func() {
		gplusGorillaMux = loadGorillaMux(gplusAPI)
	})
	calcMem("HttpRouter", func() {
		gplusHttpRouter = loadHttpRouter(gplusAPI)
	})
	calcMem("HttpTreeMux", func() {
		gplusHttpTreeMux = loadHttpTreeMux(gplusAPI)
	})
	calcMem("Macaron", func() {
		gplusMacaron = loadMacaron(gplusAPI)
	})

	println()
}

// Static
func BenchmarkBeego_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBunRouter_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBunRouter, req)
}
func BenchmarkChi_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusChi, req)
}
func BenchmarkEcho_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusEcho, req)
}
func BenchmarkFiber_GPlusStatic(b *testing.B) {
	benchFiberRequest(b, gplusFiber, "GET", "/people")
}
func BenchmarkGin_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkGojiv2_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGojiv2, req)
}
func BenchmarkGoRestful_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGoRestful, req)
}
func BenchmarkGorillaMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkMacaron_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusMacaron, req)
}

// One Param
func BenchmarkBeego_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBunRouter_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBunRouter, req)
}
func BenchmarkChi_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusChi, req)
}
func BenchmarkEcho_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusEcho, req)
}
func BenchmarkFiber_GPlusParam(b *testing.B) {
	benchFiberRequest(b, gplusFiber, "GET", "/people/118051310819094153327")
}
func BenchmarkGin_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkGojiv2_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGojiv2, req)
}
func BenchmarkGoRestful_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGoRestful, req)
}
func BenchmarkGorillaMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkMacaron_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusMacaron, req)
}

// Two Params
func BenchmarkBeego_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBunRouter_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBunRouter, req)
}
func BenchmarkChi_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusChi, req)
}
func BenchmarkEcho_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusEcho, req)
}
func BenchmarkFiber_GPlus2Params(b *testing.B) {
	benchFiberRequest(b, gplusFiber, "GET", "/people/118051310819094153327/activities/123456789")
}
func BenchmarkGin_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkGojiv2_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGojiv2, req)
}
func BenchmarkGoRestful_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGoRestful, req)
}
func BenchmarkGorillaMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkMacaron_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusMacaron, req)
}

// All Routes
func BenchmarkBeego_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBeego, gplusAPI)
}
func BenchmarkBunRouter_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBunRouter, gplusAPI)
}
func BenchmarkChi_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusChi, gplusAPI)
}
func BenchmarkEcho_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusEcho, gplusAPI)
}
func BenchmarkFiber_GPlusAll(b *testing.B) {
	benchFiberRoutes(b, gplusFiber, gplusAPI)
}
func BenchmarkGin_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGin, gplusAPI)
}
func BenchmarkGojiv2_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGojiv2, gplusAPI)
}
func BenchmarkGoRestful_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGoRestful, gplusAPI)
}
func BenchmarkGorillaMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGorillaMux, gplusAPI)
}
func BenchmarkHttpRouter_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpRouter, gplusAPI)
}
func BenchmarkHttpTreeMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpTreeMux, gplusAPI)
}
func BenchmarkMacaron_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusMacaron, gplusAPI)
}
