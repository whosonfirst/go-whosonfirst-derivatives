# go-whosonfirst-derivatives

Go package to provide a simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.

## Motivation

This is work in progress. It basically the `/api` package extracted from the [whosonfirst/go-whosonfirst-spelunker-httpd](https://github.com/whosonfirst/go-whosonfirst-spelunker-httpd).

It provides a simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.

The goal, once it's been proven to work, is to import and use the `net/http` handlers provided by this package in `whosonfirst/go-whosonfirst-spelunker-httpd`.

## Documentation

This is work in progress. Documentation is incomplete at this time.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/server cmd/server/main.go
```

### server

#### Providers

##### null://

##### reader://

#### Representations (derivative formats)

##### GeoJSON

##### GeoJSONLD

##### NavPlace

##### Select

##### Standard Places Result (SPR)

##### SVG

## See also

* https://github.com/aaronland/go-http-server
* https://github.com/sfomuseum/go-http-auth