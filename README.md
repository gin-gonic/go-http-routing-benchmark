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

Benchmark system:

- Apple M4 Pro
- go version go1.25.8 darwin/arm64
- macOS (Darwin 25.3.0)

### Summary

The table below ranks all routers by **GitHub API throughput** (203 routes, all methods), which best represents real-world routing workloads. _Lower ns/op is better._

| Rank | Router | ns/op | B/op | allocs/op | Zero-alloc |
| :--: | :----- | ----: | ---: | --------: | :--------: |
| 1 | **Gin** | 9,944 | 0 | 0 | :white_check_mark: |
| 2 | **BunRouter** | 10,281 | 0 | 0 | :white_check_mark: |
| 3 | **Echo** | 11,072 | 0 | 0 | :white_check_mark: |
| 4 | **HttpRouter** | 15,059 | 13,792 | 167 | |
| 5 | **HttpTreeMux** | 49,302 | 65,856 | 671 | |
| 6 | **Chi** | 94,376 | 130,817 | 740 | |
| 7 | **Beego** | 101,941 | 71,456 | 609 | |
| 8 | **Fiber** | 109,148 | 0 | 0 | :white_check_mark: |
| 9 | **Macaron** | 121,785 | 147,784 | 1,624 | |
| 10 | **Goji v2** | 242,849 | 313,744 | 3,712 | |
| 11 | **GoRestful** | 885,678 | 1,006,744 | 3,009 | |
| 12 | **GorillaMux** | 1,316,844 | 225,667 | 1,588 | |

**Key takeaways:**

- **Gin**, **BunRouter**, and **Echo** form the top tier — all achieve zero heap allocations and route the full GitHub API in ~10 us.
- **HttpRouter** remains extremely fast but incurs 1 alloc per parameterized route (167 allocs for 203 routes).
- **Fiber** also achieves zero allocations, but its fasthttp-based benchmark infrastructure adds per-iteration reset overhead — do not compare its absolute ns/op directly with net/http routers.
- **GorillaMux** and **GoRestful** are feature-rich but orders of magnitude slower, making them less suitable for latency-sensitive applications.

> **Fiber caveat:** Fiber benchmarks use `fasthttp.RequestCtx` with per-iteration Reset, which adds constant overhead not present in net/http benchmarks. Fiber-vs-Fiber comparisons are valid; cross-framework comparisons should be interpreted with care.

---

### Memory Consumption

The following table shows the heap memory required for loading the routing structure of the respective API.

| Router       | Static     | GitHub       | Google+    | Parse      |
| :----------- | ---------: | -----------: | ---------: | ---------: |
| HttpServeMux |   71,728 B |            — |          — |          — |
| Beego        |   98,824 B |   150,840 B  |  10,256 B  |  19,256 B  |
| BunRouter    |   51,232 B |    93,776 B  |   7,360 B  |   9,336 B  |
| Chi          |   83,160 B |    94,888 B  |   8,008 B  |   9,656 B  |
| Echo         |   91,976 B |   117,784 B  |  10,968 B  |  13,816 B  |
| Fiber        |   59,248 B |   163,832 B  |  10,840 B  |  15,352 B  |
| Gin          |   34,408 B |    58,840 B  |   4,576 B  |   7,896 B  |
| Goji v2      |  117,952 B |   118,640 B  |   8,096 B  |  16,064 B  |
| GoRestful    |  819,704 B | 1,270,848 B  |  72,536 B  | 121,200 B  |
| GorillaMux   |  599,496 B | 1,319,696 B  |  68,000 B  | 105,384 B  |
| HttpRouter   |   21,680 B |    37,072 B  |   2,776 B  |   5,024 B  |
| HttpTreeMux  |   73,448 B |    78,800 B  |   7,440 B  |   7,848 B  |
| Macaron      |   36,976 B |    90,632 B  |   8,672 B  |  13,704 B  |

### Static Routes

```
BenchmarkHttpServeMux_StaticAll   	   66117	     18172 ns/op	       0 B/op	       0 allocs/op

BenchmarkBeego_StaticAll          	   17396	     68255 ns/op	   55264 B/op	     471 allocs/op
BenchmarkBunRouter_StaticAll      	  200576	      5997 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_StaticAll            	   29058	     41317 ns/op	   57776 B/op	     314 allocs/op
BenchmarkEcho_StaticAll           	  171898	      6897 ns/op	       0 B/op	       0 allocs/op
BenchmarkFiber_StaticAll          	   40939	     29310 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_StaticAll            	  215575	      5528 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_StaticAll         	   14130	     84459 ns/op	  175840 B/op	    1099 allocs/op
BenchmarkGoRestful_StaticAll      	    2744	    436510 ns/op	  677824 B/op	    2193 allocs/op
BenchmarkGorillaMux_StaticAll     	    4200	    302825 ns/op	  133137 B/op	    1099 allocs/op
BenchmarkHttpRouter_StaticAll     	  287180	      4177 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_StaticAll    	  221372	      5363 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_StaticAll        	   14662	     81824 ns/op	  114296 B/op	    1256 allocs/op
```

### Micro Benchmarks

#### Single Route with Param

```
BenchmarkBeego_Param              	 3445593	       348.8 ns/op	     352 B/op	       3 allocs/op
BenchmarkBunRouter_Param          	100000000	        12.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_Param                	 3705072	       332.2 ns/op	     704 B/op	       4 allocs/op
BenchmarkEcho_Param               	67096422	        17.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkFiber_Param              	10718097	       114.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param                	51779376	        23.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_Param             	 2597736	       494.3 ns/op	    1136 B/op	       8 allocs/op
BenchmarkGoRestful_Param          	  760231	      1394 ns/op	    4600 B/op	      15 allocs/op
BenchmarkGorillaMux_Param         	 1979962	       630.6 ns/op	    1152 B/op	       8 allocs/op
BenchmarkHttpRouter_Param         	36953397	        31.88 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param        	 7180082	       165.0 ns/op	     352 B/op	       3 allocs/op
BenchmarkMacaron_Param            	 1768725	       708.0 ns/op	    1064 B/op	      10 allocs/op
```

#### Route with 5 Params

```
BenchmarkBeego_Param5             	 2571436	       480.3 ns/op	     352 B/op	       3 allocs/op
BenchmarkBunRouter_Param5         	29552252	        41.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_Param5               	 2662387	       453.7 ns/op	     704 B/op	       4 allocs/op
BenchmarkEcho_Param5              	27522567	        43.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkFiber_Param5             	 5022048	       271.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param5               	27722108	        44.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_Param5            	 2218150	       532.4 ns/op	    1200 B/op	       8 allocs/op
BenchmarkGoRestful_Param5         	  733239	      1579 ns/op	    4712 B/op	      15 allocs/op
BenchmarkGorillaMux_Param5        	 1236634	       972.6 ns/op	    1216 B/op	       8 allocs/op
BenchmarkHttpRouter_Param5        	13379398	        83.74 ns/op	     160 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param5       	 3437581	       358.8 ns/op	     576 B/op	       6 allocs/op
BenchmarkMacaron_Param5           	 1535364	       799.7 ns/op	    1064 B/op	      10 allocs/op
```

#### Route with 20 Params

```
BenchmarkBeego_Param20            	 1000000	      1099 ns/op	     352 B/op	       3 allocs/op
BenchmarkBunRouter_Param20        	 5585515	       211.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_Param20              	  687655	      1805 ns/op	    2504 B/op	       9 allocs/op
BenchmarkEcho_Param20             	 9540236	       127.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkFiber_Param20            	 2633816	       466.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param20              	 9690182	       121.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_Param20           	 1562407	       745.3 ns/op	    1440 B/op	       8 allocs/op
BenchmarkGoRestful_Param20        	  368145	      3337 ns/op	    7008 B/op	      20 allocs/op
BenchmarkGorillaMux_Param20       	  542091	      2223 ns/op	    3272 B/op	      13 allocs/op
BenchmarkHttpRouter_Param20       	 4168184	       290.2 ns/op	     704 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param20      	  600494	      1857 ns/op	    3144 B/op	      13 allocs/op
BenchmarkMacaron_Param20          	  493340	      2058 ns/op	    2864 B/op	      15 allocs/op
```

#### Param Write (read param + write response)

```
BenchmarkBeego_ParamWrite         	 3051739	       386.1 ns/op	     360 B/op	       4 allocs/op
BenchmarkBunRouter_ParamWrite     	47369588	        25.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_ParamWrite           	 3488102	       348.3 ns/op	     704 B/op	       4 allocs/op
BenchmarkEcho_ParamWrite          	22762137	        47.94 ns/op	       8 B/op	       1 allocs/op
BenchmarkFiber_ParamWrite         	 9702015	       125.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParamWrite           	37102452	        27.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_ParamWrite        	 2337854	       516.9 ns/op	    1168 B/op	      10 allocs/op
BenchmarkGoRestful_ParamWrite     	  803739	      1534 ns/op	    4608 B/op	      16 allocs/op
BenchmarkGorillaMux_ParamWrite    	 1899658	       665.5 ns/op	    1152 B/op	       8 allocs/op
BenchmarkHttpRouter_ParamWrite    	31189083	        37.40 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_ParamWrite   	 6040027	       180.4 ns/op	     352 B/op	       3 allocs/op
BenchmarkMacaron_ParamWrite       	 1538367	       784.3 ns/op	    1112 B/op	      13 allocs/op
```

### Parse API (26 routes)

```
BenchmarkBeego_ParseAll           	  117369	     10541 ns/op	    9152 B/op	      78 allocs/op
BenchmarkBunRouter_ParseAll       	 2039772	       588.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_ParseAll             	  142234	      8863 ns/op	   14944 B/op	      84 allocs/op
BenchmarkEcho_ParseAll            	 1603608	       742.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkFiber_ParseAll           	  279265	      4250 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseAll             	 1681222	       712.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_ParseAll          	   88848	     13264 ns/op	   29456 B/op	     199 allocs/op
BenchmarkGoRestful_ParseAll       	   23092	     54780 ns/op	  131728 B/op	     380 allocs/op
BenchmarkGorillaMux_ParseAll      	   45812	     25886 ns/op	   26960 B/op	     198 allocs/op
BenchmarkHttpRouter_ParseAll      	 1235671	       948.5 ns/op	     640 B/op	      16 allocs/op
BenchmarkHttpTreeMux_ParseAll     	  355398	      3372 ns/op	    5728 B/op	      51 allocs/op
BenchmarkMacaron_ParseAll         	   88638	     13635 ns/op	   18928 B/op	     208 allocs/op
```

### Google+ API (13 routes)

```
BenchmarkBeego_GPlusAll           	  202231	      5927 ns/op	    4576 B/op	      39 allocs/op
BenchmarkBunRouter_GPlusAll       	 3490443	       348.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_GPlusAll             	  228787	      5333 ns/op	    8480 B/op	      48 allocs/op
BenchmarkEcho_GPlusAll            	 2684996	       451.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkFiber_GPlusAll           	  474860	      2506 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusAll             	 2896033	       429.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_GPlusAll          	  146194	      8000 ns/op	   15120 B/op	     115 allocs/op
BenchmarkGoRestful_GPlusAll       	   47769	     24189 ns/op	   60720 B/op	     193 allocs/op
BenchmarkGorillaMux_GPlusAll      	   83330	     14707 ns/op	   14448 B/op	     102 allocs/op
BenchmarkHttpRouter_GPlusAll      	 1743474	       668.6 ns/op	     640 B/op	      11 allocs/op
BenchmarkHttpTreeMux_GPlusAll     	  463276	      2428 ns/op	    4032 B/op	      38 allocs/op
BenchmarkMacaron_GPlusAll         	  169158	      7294 ns/op	    9464 B/op	     104 allocs/op
```

### GitHub API (203 routes)

```
BenchmarkBeego_GithubAll          	   10000	    101941 ns/op	   71456 B/op	     609 allocs/op
BenchmarkBunRouter_GithubAll      	  117315	     10281 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_GithubAll            	   12979	     94376 ns/op	  130817 B/op	     740 allocs/op
BenchmarkEcho_GithubAll           	  106779	     11072 ns/op	       0 B/op	       0 allocs/op
BenchmarkFiber_GithubAll          	   10000	    109148 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll            	  135082	      9944 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_GithubAll         	    4814	    242849 ns/op	  313744 B/op	    3712 allocs/op
BenchmarkGoRestful_GithubAll      	    1339	    885678 ns/op	 1006744 B/op	    3009 allocs/op
BenchmarkGorillaMux_GithubAll     	     927	   1316844 ns/op	  225667 B/op	    1588 allocs/op
BenchmarkHttpRouter_GithubAll     	   82489	     15059 ns/op	   13792 B/op	     167 allocs/op
BenchmarkHttpTreeMux_GithubAll    	   24514	     49302 ns/op	   65856 B/op	     671 allocs/op
BenchmarkMacaron_GithubAll        	   10000	    121785 ns/op	  147784 B/op	    1624 allocs/op
```
