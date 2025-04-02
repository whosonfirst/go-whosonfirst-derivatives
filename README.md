# go-whosonfirst-derivatives

Go package to provide a simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.

## Motivation

This is work in progress. It basically the `/api` package extracted from the [whosonfirst/go-whosonfirst-spelunker-httpd](https://github.com/whosonfirst/go-whosonfirst-spelunker-httpd).

It provides a simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents. The goal, once it's been proven to work, is to import and use the `net/http` handlers provided by this package in `whosonfirst/go-whosonfirst-spelunker-httpd`.

## Documentation

`godoc` is incomplete at this time.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/server cmd/server/main.go
```

### server

A simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.

```
$> ./bin/server -h
A simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.

Usage:
	 ./bin/server[options]

Valid options are:
  -authenticator-uri string
    	A registered sfomuseum/go-auth.Authenticator URI. (default "null://")
  -cors-allowed-origin value
    	Zero or more allowed origins for CORS requests.
  -enable-cors
    	Enable CORS support.
  -navplace-max-features int
    	The maximum number of WOF IDs allowed in a NavPlace request. (default 10)
  -path-geojson string
    	The default path to serve GeoJSON requests from. (default "/id/{id}/geojson")
  -path-geojson-alt value
    	Zero or more alternate paths to serve GeoJSON requests from.
  -path-geojsonld string
    	The default path to serve GeoJSONLD requests from. (default "/id/{id}/geojsonld")
  -path-geojsonld-alt value
    	Zero or more alternate paths to serve GeoJSONLD requests from.
  -path-navaplace-alt value
    	Zero or more alternate paths to serve IIIF NavPlace requests from.
  -path-navplace string
    	The default path to serve IIIF NavPlace requests from. (default "/id/{id}/navplace")
  -path-select string
    	The default path to serve select requests from. (default "/id/{id}/select")
  -path-select-alt value
    	Zero or more alternate paths to serve select requests from.
  -path-spr string
    	The default path to serve standard place result (SPR) requests from. (default "/id/{id}/spr")
  -path-spr-alt value
    	Zero or more alternate paths to serve standard place result (SPR) requests from.
  -path-svg string
    	The default path to serve SVG requests from. (default "/id/{id}/svg")
  -path-svg-alt value
    	Zero or more alternate paths to serve SVG requests from.
  -provider-uri string
    	A registered whosonfirst/go-whosonfirst-derivatives.Provider URI. (default "reader://?reader-uri=https://data.whosonfirst.org")
  -server-uri string
    	A registered aaronland/go-http-server.Server URI. (default "http://localhost:8080")
  -verbose
    	Enable verbose (debug) logging.
```

#### Example

The easiest way to try things out is to run the handy `debug` Makefile target, like this:

```
$> make debug
go run -mod vendor cmd/server/main.go \
		-enable-cors \
		-cors-allowed-origin "*" \
		-verbose
2025/04/02 12:08:59 DEBUG Verbose logging enabled
2025/04/02 12:08:59 INFO Listening for requests address=http://localhost:8080
2025/04/02 12:08:59 DEBUG Enable handler uri=/id/{id}/spr handler=handler.RouteHandlerFunc
2025/04/02 12:08:59 DEBUG Enable handler uri=/id/{id}/svg handler=handler.RouteHandlerFunc
2025/04/02 12:08:59 DEBUG Enable handler uri=/id/{id}/geojson handler=handler.RouteHandlerFunc
2025/04/02 12:08:59 DEBUG Enable handler uri=/id/{id}/geojsonld handler=handler.RouteHandlerFunc
2025/04/02 12:08:59 DEBUG Enable handler uri=/id/{id}/navplace handler=handler.RouteHandlerFunc
2025/04/02 12:08:59 DEBUG Enable handler uri=/id/{id}/select handler=handler.RouteHandlerFunc
```

And then in another terminal:

```
$> curl 'http://localhost:8080/id/101736545/select?select=properties.wof:name'
"Montreal"
```

#### Representations (derivative formats)

##### GeoJSON

Returns the original Who's On First (WOF) GeoJSON document. For example `http://localhost:8080/id/101736545/geojson` would yield:

![](docs/images/go-whosonfirst-derivatives-geojson.png)

##### GeoJSONLD

Returns a Who's On First (WOF) document as a [GeoJSONLD](#) document. For example `http://localhost:8080/id/101736545/geojsonld` would yield:

![](docs/images/go-whosonfirst-derivatives-geojsonld.png)

##### NavPlace

Returns a WOF record as a GeoJSON `FeatureCollection` document. This enables WOF records to be included in [IIIF navPlace](https://preview.iiif.io/api/navplace_extension/api/extension/navplace/) records as "reference" objects. For example `http://localhost:8080/id/101736545/navplace` would yield:

![](docs/images/go-whosonfirst-derivatives-navplace.png)

You can specify multiple `Feature` records to include in a response by passing a comma-separated list of IDs. For example:

`http://localhost:8080/id/102527513,85922583,85688637/navplace`

_Note: There is a limit on the number of records that may be specified which is set by the `-navplace-max-features` flag._

##### Select

Returns a WOF record as a JSON-encoded slice of a Who's On First (WOF) GeoJSON document matching a query pattern. For example `http://localhost:8080/id/101736545/select?select=properties.wof:concordances` would yield:

![](docs/images/go-whosonfirst-derivatives-select.png)

`select` parameters should conform to the [GJSON path syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md).

As of this writing multiple `select` parameters are not supported. `select` parameters that do not match the regular expression defined in the `-select-pattern` flag (at startup) will trigger an error.

##### Standard Places Response (SPR)

Returns a WOF record as a JSON-encoded [standard places response](https://github.com/whosonfirst/go-whosonfirst-spr) (SPR) for a given WOF ID. For example `http://localhost:8080/id/101736545/spr` would yield:

![](docs/images/go-whosonfirst-derivatives-spr.png)

##### SVG

Returns a WOF record as a XML-encoded SVG representation of the geometry for a given WOF ID. For example `http://localhost:8080/id/101736545/svg` would yield:

![](docs/images/go-whosonfirst-derivatives-svg.png)

## Providers

All of the `net/http.Handler` instances used to generate derivative outputs use an instance of the `Provider` interface to retrieve a source Who's On First document to transform. That interface looks like this:

```
// Provider defines an interface for retrieving Who's On First documents used to generate derivative formats.
type Provider interface {
	// Return an `io.ReadSeekCloser` instance containing a Who's On First document.
	GetFeature(context.Context, int64, *uri.URIArgs) (io.ReadSeekCloser, error)
}
```

The following implementations of the `Provider` interface are supported by default:

### reader://

Retrieve orginal (source) Who's On First records using an implementation of the [whosonfirst/go-reader.Reader](https://github.com/whosonfirst/go-reader) interface.

The syntax for the `reader://` provider is:

```
reader://?reader-uri={REGISTERED_WHOSONFIRST_GO_READER_URI}`
```

For example:

```
reader://?reader-uri=https://data.whosonfirst.org
```

Which would retrieve orginal (source) Who's On First records from the [https://data.whosonfirst.org](https://data.whosonfirst.org) web server (using the [whosonfirst/go-reader-http](https://github.com/whosonfirst/go-reader-http) package.

### null://

This is a placeholder provider that always returns a "not found" error.

The syntax for the `null://` provider is:

```
null://
```

### Custom providers (and extending `cmd/server`)

In order to use non-default providers you will need to implement the `Provider` interface and then register it with the `RegisterProvider` method on initialization. Consult the [provider_reader.go](provider_reader.go) code for a concrete example of how to do this.

Then you will need to clone the `cmd/server/main.go` code in order to import your new package. The guts of that tool actually live in [app/server](app/server) so that these sorts of modifications can be as simple as possible. For example:

```
package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-reader-blob"	// Enable another implementation of go-reader
	_ "github.com/your-org/your-custom-provider"	// Enable your custom provider
	
	"github.com/whosonfirst/go-whosonfirst-derivatives/app/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run server, %v", err)
	}
}
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spr
* https://github.com/whosonfirst/go-whosonfirst-svg
* https://github.com/whosonfirst/go-reader
* https://github.com/whosonfirst/go-reader-http
* https://github.com/whosonfirst/go-reader-findingaid
* https://github.com/whosonfirst/go-reader-github
* https://github.com/sfomuseum/go-geojsonld
* https://github.com/aaronland/go-http-server
