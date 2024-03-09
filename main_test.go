package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kamalshkeir/kmux"
	"github.com/kamalshkeir/ksmux"
	"github.com/labstack/echo/v4"
)

// goos: windows
// goarch: amd64
// pkg: test
// cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz
// BenchmarkKmux-4                  3906224               277.9 ns/op            65 B/op          3 allocs/op
// BenchmarkFiber-4                  491118              2468 ns/op            2044 B/op         20 allocs/op
// BenchmarkKsmux-4                13126015                96.29 ns/op           18 B/op          1 allocs/op
// BenchmarkChi-4                   2875410               369.2 ns/op           353 B/op          3 allocs/op
// BenchmarkNetHTTP-4               3799891               331.4 ns/op            22 B/op          1 allocs/op
// BenchmarkGin-4                   7760228               156.8 ns/op            65 B/op          1 allocs/op
// BenchmarkEcho-4                  6918858               176.5 ns/op            27 B/op          1 allocs/op
// --- BENCH: BenchmarkEcho-4
//     main_test.go:166: --------------------------------------------------
//     main_test.go:166: --------------------------------------------------
//     main_test.go:166: --------------------------------------------------
//     main_test.go:166: --------------------------------------------------
//     main_test.go:166: --------------------------------------------------
// BenchmarkKmuxWithParam-4         3355736               363.4 ns/op            88 B/op          3 allocs/op
// BenchmarkFiberWithParam-4         472663              2450 ns/op            2058 B/op         20 allocs/op
// BenchmarkKsmuxWithParam-4        6927366               181.3 ns/op            54 B/op          1 allocs/op
// BenchmarkChiWithParam-4          2909172               390.5 ns/op           375 B/op          3 allocs/op
// BenchmarkNetHTTPWithParam-4       798625              1371 ns/op             440 B/op         11 allocs/op
// BenchmarkGinWithParam-4          5585120               216.3 ns/op            88 B/op          2 allocs/op
// BenchmarkEchoWithParam-4         5119796               220.9 ns/op            58 B/op          2 allocs/op
// --- BENCH: BenchmarkEchoWithParam-4
//     main_test.go:269: --------------------------------------------------
//     main_test.go:269: --------------------------------------------------
//     main_test.go:269: --------------------------------------------------
//     main_test.go:269: --------------------------------------------------
//     main_test.go:269: --------------------------------------------------
// BenchmarkKmuxWith2Param-4        3067424               360.9 ns/op            99 B/op          3 allocs/op
// BenchmarkFiberWith2Param-4        478682              2431 ns/op            2056 B/op         20 allocs/op
// BenchmarkKsmuxWith2Param-4       6472729               186.8 ns/op            57 B/op          1 allocs/op
// BenchmarkChiWith2Param-4         2505789               456.7 ns/op           405 B/op          3 allocs/op
// BenchmarkNetHTTPWith2Param-4     1000000              1730 ns/op             536 B/op         13 allocs/op
// BenchmarkGinWith2Param-4         4559170               261.3 ns/op           122 B/op          2 allocs/op
// BenchmarkEchoWith2Param-4        4457560               255.9 ns/op            62 B/op          2 allocs/op
// --- BENCH: BenchmarkEchoWith2Param-4
//     main_test.go:372: --------------------------------------------------
//     main_test.go:372: --------------------------------------------------
//     main_test.go:372: --------------------------------------------------
//     main_test.go:372: --------------------------------------------------
//     main_test.go:372: --------------------------------------------------
// BenchmarkKmuxWith5Param-4        2471660               472.5 ns/op           134 B/op          3 allocs/op
// BenchmarkFiberWith5Param-4        396468              2654 ns/op            2167 B/op         20 allocs/op
// BenchmarkKsmuxWith5Param-4       4053951               289.9 ns/op            98 B/op          1 allocs/op
// BenchmarkChiWith5Param-4         1725474               676.0 ns/op           445 B/op          3 allocs/op
// BenchmarkNetHTTPWith5Param-4      454626              2927 ns/op             952 B/op         17 allocs/op
// BenchmarkGinWith5Param-4         2953233               406.3 ns/op           170 B/op          2 allocs/op
// BenchmarkEchoWith5Param-4        2969812               395.8 ns/op           154 B/op          2 allocs/op
// PASS
// ok      test    44.552s

func BenchmarkKmux(b *testing.B) {
	app := kmux.New()
	app.Get("/test/bla/hello/bye", func(c *kmux.Context) {
		c.Text("Hello")
	})

	req := httptest.NewRequest("GET", "/test/bla/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}
func BenchmarkFiber(b *testing.B) {
	app := fiber.New()

	app.Get("/test/bla/hello/bye", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	req := httptest.NewRequest("GET", "/test/bla/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		adaptor.FiberApp(app).ServeHTTP(w, req)
	}
}

func BenchmarkKsmux(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/bla/hello/bye", func(c *ksmux.Context) {
		c.Text("Hello")
	})

	req := httptest.NewRequest("GET", "/test/bla/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChi(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/bla/hello/bye", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})

	req := httptest.NewRequest("GET", "/test/bla/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTP(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/bla/hello/bye", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	})

	req := httptest.NewRequest("GET", "/test/bla/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/bla/hello/bye", func(ctx *gin.Context) {
		ctx.String(200, "Hello")
	})

	req := httptest.NewRequest("GET", "/test/bla/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEcho(b *testing.B) {
	app := echo.New()
	app.GET("/test/bla/hello/bye", func(c echo.Context) error {
		return c.String(200, "Hello")
	})

	req := httptest.NewRequest("GET", "/test/bla/hello/bye", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
	b.Log("--------------------------------------------------")
}

func BenchmarkKmuxWithParam(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:something", func(c *kmux.Context) {
		c.Text("Hello " + c.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFiberWithParam(b *testing.B) {
	app := fiber.New()

	app.Get("/test/:something", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		adaptor.FiberApp(app).ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWithParam(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:something", func(c *ksmux.Context) {
		c.Text("Hello " + c.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChiWithParam(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/{something}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello " + chi.URLParam(r, "something")))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWithParam(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{something}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello " + r.PathValue("something")))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWithParam(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:something", func(ctx *gin.Context) {
		ctx.String(200, "Hello "+ctx.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWithParam(b *testing.B) {
	app := echo.New()
	app.GET("/test/:something", func(c echo.Context) error {
		return c.String(200, "Hello "+c.Param("something"))
	})

	req := httptest.NewRequest("GET", "/test/user1", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
	b.Log("--------------------------------------------------")
}

func BenchmarkKmuxWith2Param(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:something/:another", func(c *kmux.Context) {
		c.Text("Hello " + c.Param("something") + c.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFiberWith2Param(b *testing.B) {
	app := fiber.New()

	app.Get("/test/:something/:another", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("something") + c.Params("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		adaptor.FiberApp(app).ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWith2Param(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:something/:another", func(c *ksmux.Context) {
		c.Text("Hello " + c.Param("something") + c.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChiWith2Param(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/{something}/{another}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello " + chi.URLParam(r, "something") + chi.URLParam(r, "another")))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWith2Param(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{something}/{another}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello, " + r.PathValue("something") + r.PathValue("another")))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWith2Param(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:something/:another", func(ctx *gin.Context) {
		ctx.String(200, "Hello "+ctx.Param("something")+ctx.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWith2Param(b *testing.B) {
	app := echo.New()
	app.GET("/test/:something/:another", func(c echo.Context) error {
		return c.String(200, "Hello "+c.Param("something")+c.Param("another"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
	b.Log("--------------------------------------------------")
}

func BenchmarkKmuxWith5Param(b *testing.B) {
	app := kmux.New()
	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c *kmux.Context) {
		c.Text("Hello " + c.Param("first") + c.Param("second") + c.Param("third") + c.Param("fourth") + c.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkFiberWith5Param(b *testing.B) {
	app := fiber.New()

	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("first") + c.Params("second") + c.Params("third") + c.Params("fourth") + c.Params("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		adaptor.FiberApp(app).ServeHTTP(w, req)
	}
}

func BenchmarkKsmuxWith5Param(b *testing.B) {
	app := ksmux.New()
	app.Get("/test/:first/:second/:third/:fourth/:fifth", func(c *ksmux.Context) {
		c.Text("Hello " + c.Param("first") + c.Param("second") + c.Param("third") + c.Param("fourth") + c.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkChiWith5Param(b *testing.B) {
	app := chi.NewRouter()
	app.Get("/test/{first}/{second}/{third}/{fourth}/{fifth}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello " + chi.URLParam(r, "first") + chi.URLParam(r, "second") + chi.URLParam(r, "third") + chi.URLParam(r, "fourth") + chi.URLParam(r, "fifth")))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkNetHTTPWith5Param(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test/{first}/{second}/{third}/{fourth}/{fifth}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "Hello, "+r.PathValue("first")+r.PathValue("second")+r.PathValue("third")+r.PathValue("fourth")+r.PathValue("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}

func BenchmarkGinWith5Param(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.GET("/test/:first/:second/:third/:fourth/:fifth", func(ctx *gin.Context) {
		ctx.String(200, "Hello "+ctx.Param("first")+ctx.Param("second")+ctx.Param("third")+ctx.Param("fourth")+ctx.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}

func BenchmarkEchoWith5Param(b *testing.B) {
	app := echo.New()
	app.GET("/test/:first/:second/:third/:fourth/:fifth", func(c echo.Context) error {
		return c.String(200, "Hello "+c.Param("first")+c.Param("second")+c.Param("third")+c.Param("fourth")+c.Param("fifth"))
	})

	req := httptest.NewRequest("GET", "/test/user1/more/one/two/three", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, req)
	}
}
