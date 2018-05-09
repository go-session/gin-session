# Session middleware for [Gin](https://github.com/gin-gonic/gin)

[![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

## Quick Start

### Download and install

```bash
$ go get -u -v github.com/go-session/gin-session
```

### Create file `server.go`

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
	"gopkg.in/session.v2"
)

func main() {
	app := gin.Default()

	app.Use(ginsession.New(
		session.SetCookieName("session_id"),
		session.SetSign([]byte("sign")),
	))

	app.GET("/", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		store.Set("foo", "bar")
		err := store.Save()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.Redirect(302, "/foo")
	})

	app.GET("/foo", func(ctx *gin.Context) {
		store := ginsession.FromContext(ctx)
		foo, ok := store.Get("foo")
		if !ok {
			ctx.AbortWithStatus(404)
			return
		}
		ctx.String(http.StatusOK, "foo:%s", foo)
	})

	app.Run(":8080")
}
```

### Build and run

```bash
$ go build server.go
$ ./server
```

### Open in your web browser

<http://localhost:8080>

    foo:bar


## MIT License

    Copyright (c) 2018 Lyric

[reportcard-url]: https://goreportcard.com/report/github.com/go-session/gin-session
[reportcard-image]: https://goreportcard.com/badge/github.com/go-session/gin-session
[godoc-url]: https://godoc.org/github.com/go-session/gin-session
[godoc-image]: https://godoc.org/github.com/go-session/gin-session?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg