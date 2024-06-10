package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
	const (
		authenticationSrvAddr = "http://localhost:1234/login"
	)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/login", forward(authenticationSrvAddr))
	http.ListenAndServe(":3000", r)
}

func forward(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(r.Method, url, r.Body)
		if err != nil {
			logger.With("error", err).With("url", url).Error("failed to create new request")
			http.Error(w, http.StatusText(503), 503)
			return
		}

		req.Header = r.Header
		req.URL.RawQuery = r.URL.RawQuery

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.With("error", err).With("url", url).Error("") // dont know what message to log yet!
			http.Error(w, http.StatusText(503), 503)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}
