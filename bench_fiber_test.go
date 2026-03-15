// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func benchFiberRequest(b *testing.B, app *fiber.App, method, path string) {
	handler := app.Handler()
	ctx := &fasthttp.RequestCtx{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx.Request.Reset()
		ctx.Response.Reset()
		ctx.Request.Header.SetMethod(method)
		ctx.Request.SetRequestURI(path)
		handler(ctx)
	}
}

func benchFiberRoutes(b *testing.B, app *fiber.App, routes []route) {
	handler := app.Handler()
	ctx := &fasthttp.RequestCtx{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, r := range routes {
			ctx.Request.Reset()
			ctx.Response.Reset()
			ctx.Request.Header.SetMethod(r.method)
			ctx.Request.SetRequestURI(r.path)
			handler(ctx)
		}
	}
}

// TestFiberRouters validates Fiber routing correctness.
// Fiber uses fasthttp and cannot participate in the net/http TestRouters test.
func TestFiberRouters(t *testing.T) {
	loadTestHandler = true
	defer func() { loadTestHandler = false }()

	for _, api := range apis {
		app := loadFiber(api.routes)
		handler := app.Handler()

		for _, route := range api.routes {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.Header.SetMethod(route.method)
			ctx.Request.SetRequestURI(route.path)
			handler(ctx)

			code := ctx.Response.StatusCode()
			body := string(ctx.Response.Body())
			if code != 200 || body != route.path {
				t.Errorf(
					"Fiber in API %s: %d - %s; expected %s %s\n",
					api.name, code, body, route.method, route.path,
				)
			}
		}
	}
}

// Micro Benchmarks

// Route with Param (no write)
func BenchmarkFiber_Param(b *testing.B) {
	app := loadFiberSingle("GET", "/user/:name", fiberHandler)
	benchFiberRequest(b, app, "GET", "/user/gordon")
}

// Route with 5 Params (no write)
func BenchmarkFiber_Param5(b *testing.B) {
	app := loadFiberSingle("GET", fiveColon, fiberHandler)
	benchFiberRequest(b, app, "GET", fiveRoute)
}

// Route with 20 Params (no write)
func BenchmarkFiber_Param20(b *testing.B) {
	app := loadFiberSingle("GET", twentyColon, fiberHandler)
	benchFiberRequest(b, app, "GET", twentyRoute)
}

// Route with Param and write
func BenchmarkFiber_ParamWrite(b *testing.B) {
	app := loadFiberSingle("GET", "/user/:name", fiberHandlerWrite)
	benchFiberRequest(b, app, "GET", "/user/gordon")
}
