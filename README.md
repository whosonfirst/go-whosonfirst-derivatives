# go-whosonfirst-derivatives

This is work in progress. It basically the `/api` package extracted from the [whosonfirst/go-whosonfirst-spelunker-httpd](https://github.com/whosonfirst/go-whosonfirst-spelunker-httpd).

It provides a simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.

The goal, once it's been proven to work, is to import and use the `net/http` handlers provided by this package in `whosonfirst/go-whosonfirst-spelunker-httpd`.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/server cmd/server/main.go
```

### server

#### Providers

#### Representations (derivative formats)

## See also

* 