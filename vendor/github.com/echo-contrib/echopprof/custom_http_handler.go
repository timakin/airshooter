package echopprof

import "net/http"

type customHTTPHandler struct {
	serveHTTP func(w http.ResponseWriter, r *http.Request)
}

func (c *customHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.serveHTTP(w, r)
}
