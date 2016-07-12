# echopprof
A wrapper for golang web framework echo to use net/http/pprof easily.
# install
First install echopprof to your GOPATH using go get:
```
go get github.com/echo-contrib/echopprof
```
# Usage
```
package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/echo-contrib/echopprof"
)

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "Hello, World!\n")
	}
}

func main() {
    e := echo.New()
    e.Get("/", hello())

    // automatically add routers for net/http/pprof
    // e.g. /debug/pprof, /debug/pprof/heap, etc.
    echopprof.Wrapper(e)
    e.Run(standard.New(":8080"))
}
```
Start this server, and now visit http://127.0.0.1:8080/debug/pprof/ and you'll see what you want.



