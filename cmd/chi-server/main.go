package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"net/http/pprof"
	"time"
)

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", "123")

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func startWelcome() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(MyMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/foo", func(w http.ResponseWriter, r *http.Request) {
		testFlag := fmt.Sprintf("%v", r.Context().Value("testFlag"))
		s := make([]string, 3)
		for i:= 0; i < 10000000; i++ {
			s = append(s, "magical pandas")
			if (i % 100000) == 0 {
				time.Sleep(500 * time.Millisecond)
			}
		}
		w.Write([]byte(testFlag))
	})

	r.HandleFunc("/debug/pprof/*", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":3000", r)
}

func main() {
	startWelcome()
}
