package echopprof

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/engine/fasthttp"
	"net/http"
	"log"
	"sync"
)

type customEchoHandler struct {
	httpHandler http.Handler

	wrappedHandleFunc echo.HandlerFunc
	once sync.Once
}

func (ceh *customEchoHandler) Handle(c echo.Context) error {
	ceh.once.Do(func() {
		ceh.wrappedHandleFunc = ceh.mustWrapHandleFunc(c)
	})
	return ceh.wrappedHandleFunc(c)
}

func (ceh *customEchoHandler) mustWrapHandleFunc(c echo.Context) echo.HandlerFunc {
	if _, ok := c.Request().(*standard.Request); ok {
		return standard.WrapHandler(ceh.httpHandler)
	} else if _, ok = c.Request().(*fasthttp.Request); ok {
		return NewFastHTTPEchoAdaptor(ceh.httpHandler)
	}

	log.Fatal("Unknown HTTP implementation")
	return nil
}

func fromHTTPHandler(httpHandler http.Handler) *customEchoHandler {
	return &customEchoHandler{ httpHandler: httpHandler }
}

func fromHandlerFunc(serveHTTP func(w http.ResponseWriter, r *http.Request)) *customEchoHandler {
	return &customEchoHandler{ httpHandler: &customHTTPHandler{ serveHTTP: serveHTTP }}
}
