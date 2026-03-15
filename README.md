# Go HTTP Router Benchmark

This benchmark suite aims to compare the performance of HTTP request routers for [Go](https://golang.org) by implementing the routing structure of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.

Of course the tested routers can be used for any kind of HTTP request → handler function routing, not only (REST) APIs.

## Tested routers & frameworks

- [Beego](http://beego.me/)
- [BunRouter](https://github.com/uptrace/bunrouter)
- [Chi](https://github.com/go-chi/chi)
- [Echo](https://github.com/labstack/echo)
- [Fiber](https://github.com/gofiber/fiber) _(fasthttp-based, benchmarked separately)_
- [Gin](https://github.com/gin-gonic/gin)
- [Goji v2](https://github.com/goji/goji)
- [Go-Restful](https://github.com/emicklei/go-restful)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [http.ServeMux](https://pkg.go.dev/net/http#ServeMux) _(static routes only)_
- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [HttpTreeMux](https://github.com/dimfeld/httptreemux)
- [Macaron](https://github.com/go-macaron/macaron)

> **Note:** Fiber is based on [fasthttp](https://github.com/valyala/fasthttp) and does not implement the `net/http` `Handler` interface. Its benchmarks use a separate infrastructure and should be compared independently from the `net/http`-based routers.

## Motivation

Go is a great language for web applications. Since the [default _request multiplexer_](https://pkg.go.dev/net/http#ServeMux) of Go's net/http package is very simple and limited, an accordingly high number of HTTP request routers exist.

Unfortunately, most of the (early) routers use pretty bad routing algorithms. Moreover, many of them are very wasteful with memory allocations, which can become a problem in a language with Garbage Collection like Go, since every (heap) allocation results in more work for the Garbage Collector.

Lately more and more bloated frameworks pop up, outdoing one another in the number of features. This benchmark tries to measure their overhead.

Be aware that we are comparing apples and oranges here. We compare feature-rich frameworks to packages with simple routing functionality only. But since we are only interested in decent request routing, I think this is not entirely unfair. The frameworks are configured to do as little additional work as possible.

If you care about performance, this benchmark can maybe help you find the right router, which scales with your application.

Personally, I prefer slim and optimized software, which is why I implemented [HttpRouter](https://github.com/julienschmidt/httprouter), which is also tested here. In fact, this benchmark suite started as part of the packages tests, but was then extended to a generic benchmark suite.
So keep in mind, that I am not completely unbiased :relieved:

## Usage

```sh
go test -bench=. -benchmem -timeout=20m
```

## Results

_Benchmark results have been cleared after the framework cleanup. Run the benchmarks on your own machine to generate fresh results._
