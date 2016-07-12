package echopprof

import (
	"github.com/labstack/echo"
	"net/http/pprof"
)

func Wrap(e *echo.Echo) {
	e.GET("/debug/pprof/", fromHandlerFunc(pprof.Index).Handle)
	e.GET("/debug/pprof/heap", fromHTTPHandler(pprof.Handler("heap")).Handle)
	e.GET("/debug/pprof/goroutine", fromHTTPHandler(pprof.Handler("goroutine")).Handle)
	e.GET("/debug/pprof/block", fromHTTPHandler(pprof.Handler("block")).Handle)
	e.GET("/debug/pprof/threadcreate", fromHTTPHandler(pprof.Handler("threadcreate")).Handle)
	e.GET("/debug/pprof/cmdline", fromHandlerFunc(pprof.Cmdline).Handle)
	e.GET("/debug/pprof/profile", fromHandlerFunc(pprof.Profile).Handle)
	e.GET("/debug/pprof/symbol", fromHandlerFunc(pprof.Symbol).Handle)
}

var Wrapper = Wrap
