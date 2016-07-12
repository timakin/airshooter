package echopprof
import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"net/url"
	"io"
	"fmt"
	"strings"
)

const textPlainContentType = "text/plain"

// NewFastHTTPEchoAdaptor is responsible for adapting echo requests through fasthttp interfaces to net/http requests.
//
// Based on valyala/fasthttp implementation.
// Available here: https://github.com/valyala/fasthttp/blob/master/fasthttpadaptor/adaptor.go
func NewFastHTTPEchoAdaptor(h http.Handler) echo.HandlerFunc  {
	return func(c echo.Context) error {
		var r http.Request
		ctx := c.Request().(*fasthttp.Request).RequestCtx

		body := ctx.PostBody()
		r.Method = string(ctx.Method())
		r.Proto = "HTTP/1.1"
		r.ProtoMajor = 1
		r.ProtoMinor = 1
		r.RequestURI = string(ctx.RequestURI())
		r.ContentLength = int64(len(body))
		r.Host = string(ctx.Host())
		r.RemoteAddr = ctx.RemoteAddr().String()

		hdr := make(http.Header)
		ctx.Request.Header.VisitAll(func(k, v []byte) {
			hdr.Set(string(k), string(v))
		})
		r.Header = hdr
		r.Body = &netHTTPBody{body}
		rURL, err := url.ParseRequestURI(r.RequestURI)
		if err != nil {
			ctx.Logger().Printf("cannot parse requestURI %q: %s", r.RequestURI, err)
			return fmt.Errorf("Internal Server Error")
		}
		r.URL = rURL

		var w netHTTPResponseWriter
		h.ServeHTTP(&w, &r)

		ctx.SetStatusCode(w.StatusCode())
		for k, vv := range w.Header() {
			for _, v := range vv {
				c.Response().Header().Set(k, v)
			}
		}

		if strings.Contains(c.Response().Header().Get(echo.HeaderContentType), textPlainContentType) {
			c.Response().Header().Set(echo.HeaderContentType, http.DetectContentType(w.body))
		}
		c.Response().Write(w.body)
		return nil
	}
}

type netHTTPBody struct {
	b []byte
}

func (r *netHTTPBody) Read(p []byte) (int, error) {
	if len(r.b) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.b)
	r.b = r.b[n:]
	return n, nil
}

func (r *netHTTPBody) Close() error {
	r.b = r.b[:0]
	return nil
}

type netHTTPResponseWriter struct {
	statusCode int
	h          http.Header
	body       []byte
}

func (w *netHTTPResponseWriter) StatusCode() int {
	if w.statusCode == 0 {
		return http.StatusOK
	}
	return w.statusCode
}

func (w *netHTTPResponseWriter) Header() http.Header {
	if w.h == nil {
		w.h = make(http.Header)
	}
	return w.h
}

func (w *netHTTPResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *netHTTPResponseWriter) Write(p []byte) (int, error) {
	w.body = append(w.body, p...)
	return len(p), nil
}
