package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/kamalshkeir/kmux"
	"github.com/kamalshkeir/ksmux"
	"github.com/labstack/echo/v4"
)

// goos: windows
// goarch: amd64
// pkg: test
// cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz
// BenchmarkKmux-4                  4367367               264.2 ns/op            63 B/op          3 allocs/op
// BenchmarkKsmux-4                12932511                88.70 ns/op           18 B/op          1 allocs/op
// BenchmarkChi-4                   3292162               310.6 ns/op           351 B/op          3 allocs/op
// BenchmarkNetHTTP-4               3800149               319.6 ns/op            22 B/op          1 allocs/op
// BenchmarkGin-4                   7743439               144.9 ns/op            65 B/op          1 allocs/op
// BenchmarkEcho-4                  6710234               174.5 ns/op            18 B/op          1 allocs/op
// --- BENCH: BenchmarkEcho-4
//     main_test.go:142: --------------------------------------------------
//     main_test.go:142: --------------------------------------------------
//     main_test.go:142: --------------------------------------------------
//     main_test.go:142: --------------------------------------------------
//     main_test.go:142: --------------------------------------------------
// BenchmarkKmuxWithParam-4         3572722               334.2 ns/op            85 B/op          3 allocs/op
// BenchmarkKsmuxWithParam-4        7835623               155.8 ns/op            50 B/op          1 allocs/op
// BenchmarkChiWithParam-4          2796444               383.8 ns/op           376 B/op          3 allocs/op
// BenchmarkNetHTTPWithParam-4       767793              1363 ns/op             440 B/op         11 allocs/op
// BenchmarkGinWithParam-4          5606192               209.8 ns/op            87 B/op          2 allocs/op
// BenchmarkEchoWithParam-4         5107010               224.6 ns/op            58 B/op          2 allocs/op
// --- BENCH: BenchmarkEchoWithParam-4
//     main_test.go:230: --------------------------------------------------
//     main_test.go:230: --------------------------------------------------
//     main_test.go:230: --------------------------------------------------
//     main_test.go:230: --------------------------------------------------
//     main_test.go:230: --------------------------------------------------
// BenchmarkKmuxWith2Param-4        3360949               352.9 ns/op            95 B/op          3 allocs/op
// BenchmarkKsmuxWith2Param-4       6263086               186.8 ns/op            58 B/op          1 allocs/op
// BenchmarkNetHTTPWith2Param-4     1000000              1694 ns/op             536 B/op         13 allocs/op
// BenchmarkGinWith2Param-4         5050681               250.3 ns/op           117 B/op          2 allocs/op
// BenchmarkEchoWith2Param-4        4527847               265.8 ns/op            91 B/op          2 allocs/op
// --- BENCH: BenchmarkEchoWith2Param-4
//     main_test.go:303: --------------------------------------------------
//     main_test.go:303: --------------------------------------------------
//     main_test.go:303: --------------------------------------------------
//     main_test.go:303: --------------------------------------------------
//     main_test.go:303: --------------------------------------------------
// BenchmarkKmuxWith5Param-4        2736858               419.9 ns/op            88 B/op          3 allocs/op
// BenchmarkKsmuxWith5Param-4       3952650               292.5 ns/op            99 B/op          1 allocs/op
// BenchmarkNetHTTPWith5Param-4      405996              2594 ns/op             952 B/op         17 allocs/op
// BenchmarkGinWith5Param-4         3387904               367.1 ns/op           159 B/op          2 allocs/op
// BenchmarkEchoWith5Param-4        2962162               402.5 ns/op           154 B/op          2 allocs/op
// PASS
// ok      test    31.733s

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
