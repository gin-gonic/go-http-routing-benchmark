// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"

	// If you add new routers please:
	// - Keep the benchmark functions etc. alphabetically sorted
	// - Make a pull request (without benchmark results) at
	//   https://github.com/julienschmidt/go-http-routing-benchmark
	"github.com/0mjs/zinc"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dimfeld/httptreemux"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bunrouter"
	gojiv2 "goji.io"
	gojiv2pat "goji.io/pat"
	"gopkg.in/macaron.v1"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

var nullLogger *log.Logger

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initBeego()
	initGin()
}

// Common
func httpHandlerFunc(_ http.ResponseWriter, _ *http.Request) {}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

// Beego
func beegoHandler(ctx *context.Context) {}

func beegoHandlerWrite(ctx *context.Context) {
	ctx.WriteString(ctx.Input.Param(":name"))
}

func beegoHandlerTest(ctx *context.Context) {
	ctx.WriteString(ctx.Request.RequestURI)
}

func initBeego() {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
}

func loadBeego(routes []route) http.Handler {
	h := beegoHandler
	if loadTestHandler {
		h = beegoHandlerTest
	}

	re := regexp.MustCompile(":([^/]*)")
	app := beego.NewControllerRegister()
	for _, route := range routes {
		route.path = re.ReplaceAllString(route.path, ":$1")
		switch route.method {
		case "GET":
			app.Get(route.path, h)
		case "POST":
			app.Post(route.path, h)
		case "PUT":
			app.Put(route.path, h)
		case "PATCH":
			app.Patch(route.path, h)
		case "DELETE":
			app.Delete(route.path, h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return app
}

func loadBeegoSingle(method, path string, handler beego.FilterFunc) http.Handler {
	app := beego.NewControllerRegister()
	switch method {
	case "GET":
		app.Get(path, handler)
	case "POST":
		app.Post(path, handler)
	case "PUT":
		app.Put(path, handler)
	case "PATCH":
		app.Patch(path, handler)
	case "DELETE":
		app.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return app
}

// BunRouter
func bunRouterHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func bunRouterHandlerWrite(w http.ResponseWriter, req bunrouter.Request) error {
	io.WriteString(w, req.Param("name"))
	return nil
}

func bunRouterHandlerTest(w http.ResponseWriter, req bunrouter.Request) error {
	io.WriteString(w, req.RequestURI)
	return nil
}

func loadBunRouter(routes []route) http.Handler {
	h := bunRouterHandler
	if loadTestHandler {
		h = bunRouterHandlerTest
	}

	router := bunrouter.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.GET(route.path, h)
		case "POST":
			router.POST(route.path, h)
		case "PUT":
			router.PUT(route.path, h)
		case "PATCH":
			router.PATCH(route.path, h)
		case "DELETE":
			router.DELETE(route.path, h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return router
}

func loadBunRouterSingle(method, path string, handler bunrouter.HandlerFunc) http.Handler {
	router := bunrouter.New()
	switch method {
	case "GET":
		router.GET(path, handler)
	case "POST":
		router.POST(path, handler)
	case "PUT":
		router.PUT(path, handler)
	case "PATCH":
		router.PATCH(path, handler)
	case "DELETE":
		router.DELETE(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return router
}

// Chi
func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, chi.URLParam(r, "name"))
}

func loadChi(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")

	mux := chi.NewRouter()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "{$1}")

		switch route.method {
		case "GET":
			mux.Get(path, h)
		case "POST":
			mux.Post(path, h)
		case "PUT":
			mux.Put(path, h)
		case "PATCH":
			mux.Patch(path, h)
		case "DELETE":
			mux.Delete(path, h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return mux
}

func loadChiSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := chi.NewRouter()
	switch method {
	case "GET":
		mux.Get(path, handler)
	case "POST":
		mux.Post(path, handler)
	case "PUT":
		mux.Put(path, handler)
	case "PATCH":
		mux.Patch(path, handler)
	case "DELETE":
		mux.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return mux
}

// Echo
func echoHandler(c echo.Context) error {
	return nil
}

func echoHandlerWrite(c echo.Context) error {
	io.WriteString(c.Response(), c.Param("name"))
	return nil
}

func echoHandlerTest(c echo.Context) error {
	io.WriteString(c.Response(), c.Request().RequestURI)
	return nil
}

func loadEcho(routes []route) http.Handler {
	var h echo.HandlerFunc = echoHandler
	if loadTestHandler {
		h = echoHandlerTest
	}

	e := echo.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			e.GET(r.path, h)
		case "POST":
			e.POST(r.path, h)
		case "PUT":
			e.PUT(r.path, h)
		case "PATCH":
			e.PATCH(r.path, h)
		case "DELETE":
			e.DELETE(r.path, h)
		default:
			panic("Unknown HTTP method: " + r.method)
		}
	}
	return e
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	case "POST":
		e.POST(path, h)
	case "PUT":
		e.PUT(path, h)
	case "PATCH":
		e.PATCH(path, h)
	case "DELETE":
		e.DELETE(path, h)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return e
}

// Fiber (fasthttp-based, not net/http compatible)
func fiberHandler(c *fiber.Ctx) error {
	return nil
}

func fiberHandlerWrite(c *fiber.Ctx) error {
	_, _ = c.WriteString(c.Params("name"))
	return nil
}

func fiberHandlerTest(c *fiber.Ctx) error {
	_, _ = c.WriteString(c.Path())
	return nil
}

func loadFiber(routes []route) *fiber.App {
	var h fiber.Handler = fiberHandler
	if loadTestHandler {
		h = fiberHandlerTest
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	for _, r := range routes {
		app.Add(r.method, r.path, h)
	}
	return app
}

func loadFiberSingle(method, path string, handler fiber.Handler) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Add(method, path, handler)
	return app
}

// Gin
func ginHandle(_ *gin.Context) {}

func ginHandleWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Params.ByName("name"))
}

func ginHandleTest(c *gin.Context) {
	io.WriteString(c.Writer, c.Request.RequestURI)
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []route) http.Handler {
	h := ginHandle
	if loadTestHandler {
		h = ginHandleTest
	}

	router := gin.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, handle)
	return router
}

// Goji v2 (goji.io)
func gojiv2Handler(w http.ResponseWriter, r *http.Request) {}

func gojiv2HandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, gojiv2pat.Param(r, "name"))
}

func gojiv2HandlerTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

func loadGojiv2(routes []route) http.Handler {
	h := gojiv2Handler
	if loadTestHandler {
		h = gojiv2HandlerTest
	}

	mux := gojiv2.NewMux()
	for _, route := range routes {
		switch route.method {
		case "GET":
			mux.HandleFunc(gojiv2pat.Get(route.path), h)
		case "POST":
			mux.HandleFunc(gojiv2pat.Post(route.path), h)
		case "PUT":
			mux.HandleFunc(gojiv2pat.Put(route.path), h)
		case "PATCH":
			mux.HandleFunc(gojiv2pat.Patch(route.path), h)
		case "DELETE":
			mux.HandleFunc(gojiv2pat.Delete(route.path), h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return mux
}

func loadGojiv2Single(method, path string, handler func(http.ResponseWriter, *http.Request)) http.Handler {
	mux := gojiv2.NewMux()
	switch method {
	case "GET":
		mux.HandleFunc(gojiv2pat.Get(path), handler)
	case "POST":
		mux.HandleFunc(gojiv2pat.Post(path), handler)
	case "PUT":
		mux.HandleFunc(gojiv2pat.Put(path), handler)
	case "PATCH":
		mux.HandleFunc(gojiv2pat.Patch(path), handler)
	case "DELETE":
		mux.HandleFunc(gojiv2pat.Delete(path), handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return mux
}

// go-restful
func goRestfulHandler(r *restful.Request, w *restful.Response) {}

func goRestfulHandlerWrite(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.PathParameter("name"))
}

func goRestfulHandlerTest(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.Request.RequestURI)
}

func loadGoRestful(routes []route) http.Handler {
	h := goRestfulHandler
	if loadTestHandler {
		h = goRestfulHandlerTest
	}

	re := regexp.MustCompile(":([^/]*)")

	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)

	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "{$1}")

		switch route.method {
		case "GET":
			ws.Route(ws.GET(path).To(h))
		case "POST":
			ws.Route(ws.POST(path).To(h))
		case "PUT":
			ws.Route(ws.PUT(path).To(h))
		case "PATCH":
			ws.Route(ws.PATCH(path).To(h))
		case "DELETE":
			ws.Route(ws.DELETE(path).To(h))
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	wsContainer.Add(ws)
	return wsContainer
}

func loadGoRestfulSingle(method, path string, handler restful.RouteFunction) http.Handler {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)
	switch method {
	case "GET":
		ws.Route(ws.GET(path).To(handler))
	case "POST":
		ws.Route(ws.POST(path).To(handler))
	case "PUT":
		ws.Route(ws.PUT(path).To(handler))
	case "PATCH":
		ws.Route(ws.PATCH(path).To(handler))
	case "DELETE":
		ws.Route(ws.DELETE(path).To(handler))
	default:
		panic("Unknown HTTP method: " + method)
	}
	wsContainer.Add(ws)
	return wsContainer
}

// gorilla/mux
func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params["name"])
}

func loadGorillaMux(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.path, "{$1}"),
			h,
		).Methods(route.method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// HttpRouter
func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

func httpRouterHandleWrite(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func httpRouterHandleTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpRouter(routes []route) http.Handler {
	h := httpRouterHandle
	if loadTestHandler {
		h = httpRouterHandleTest
	}

	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

// httpTreeMux
func httpTreeMuxHandler(_ http.ResponseWriter, _ *http.Request, _ map[string]string) {}

func httpTreeMuxHandlerWrite(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func httpTreeMuxHandlerTest(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpTreeMux(routes []route) http.Handler {
	h := httpTreeMuxHandler
	if loadTestHandler {
		h = httpTreeMuxHandlerTest
	}

	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpTreeMuxSingle(method, path string, handler httptreemux.HandlerFunc) http.Handler {
	router := httptreemux.New()
	router.Handle(method, path, handler)
	return router
}

// Macaron
func macaronHandler() {}

func macaronHandlerWrite(c *macaron.Context) string {
	return c.Params("name")
}

func macaronHandlerTest(c *macaron.Context) string {
	return c.Req.RequestURI
}

func loadMacaron(routes []route) http.Handler {
	var h = []macaron.Handler{macaronHandler}
	if loadTestHandler {
		h[0] = macaronHandlerTest
	}

	m := macaron.New()
	for _, route := range routes {
		m.Handle(route.method, route.path, h)
	}
	return m
}

func loadMacaronSingle(method, path string, handler interface{}) http.Handler {
	m := macaron.New()
	m.Handle(method, path, []macaron.Handler{handler})
	return m
}

// Zinc
func zincHandler(_ *zinc.Context) error {
	return nil
}

func zincHandlerWrite(c *zinc.Context) error {
	_, err := io.WriteString(c.Writer(), c.Param("name"))
	return err
}

func zincHandlerTest(c *zinc.Context) error {
	_, err := io.WriteString(c.Writer(), c.Request().RequestURI)
	return err
}

func zincPath(path string) string {
	if !strings.ContainsAny(path, ":*") {
		return path
	}

	segments := strings.Split(path, "/")
	for i, segment := range segments {
		switch {
		case strings.HasPrefix(segment, ":"):
			segments[i] = "{" + segment[1:] + "}"
		case strings.HasPrefix(segment, "*"):
			segments[i] = "{" + segment[1:] + "...}"
		}
	}
	return strings.Join(segments, "/")
}

func loadZinc(routes []route) http.Handler {
	h := zincHandler
	if loadTestHandler {
		h = zincHandlerTest
	}

	app := zinc.New()
	for _, route := range routes {
		app.Add(route.method, zincPath(route.path), h)
	}
	return app
}

func loadZincSingle(method, path string, handler zinc.HandlerFunc) http.Handler {
	app := zinc.New()
	app.Add(method, zincPath(path), handler)
	return app
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
