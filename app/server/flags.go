package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var server_uri string
var provider_uri string

var path_geojson string
var path_geojson_alt multi.MultiCSVString

var path_geojsonld string
var path_geojsonld_alt multi.MultiCSVString

var path_navplace string
var path_navplace_alt multi.MultiCSVString

var path_select string
var path_select_alt multi.MultiCSVString

var path_spr string
var path_spr_alt multi.MultiCSVString

var path_svg string
var path_svg_alt multi.MultiCSVString

var navplace_max_features int

var enable_cors bool
var cors_allowed_origins multi.MultiCSVString

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("server")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A registered aaronland/go-http-server.Server URI.")
	fs.StringVar(&provider_uri, "provider-uri", "reader://?reader-uri=https://data.whosonfirst.org", "A registered whosonfirst/go-whosonfirst-derivatives.Provider URI.")

	fs.StringVar(&path_geojson, "path-geojson", "/id/{id}/geojson", "The default path to serve GeoJSON requests from.")
	fs.Var(&path_geojson_alt, "path-geojson-alt", "Zero or more alternate paths to serve GeoJSON requests from.")

	fs.StringVar(&path_geojsonld, "path-geojsonld", "/id/{id}/geojsonld", "The default path to serve GeoJSONLD requests from.")
	fs.Var(&path_geojsonld_alt, "path-geojsonld-alt", "Zero or more alternate paths to serve GeoJSONLD requests from.")

	fs.StringVar(&path_navplace, "path-navplace", "/id/{id}/navplace", "The default path to serve IIIF NavPlace requests from.")
	fs.Var(&path_navplace_alt, "path-navaplace-alt", "Zero or more alternate paths to serve IIIF NavPlace requests from.")

	fs.StringVar(&path_select, "path-select", "/id/{id}/select", "The default path to serve select requests from.")
	fs.Var(&path_select_alt, "path-select-alt", "Zero or more alternate paths to serve select requests from.")

	fs.StringVar(&path_spr, "path-spr", "/id/{id}/spr", "The default path to serve standard place result (SPR) requests from.")
	fs.Var(&path_spr_alt, "path-spr-alt", "Zero or more alternate paths to serve standard place result (SPR) requests from.")

	fs.StringVar(&path_svg, "path-svg", "/id/{id}/svg", "The default path to serve SVG requests from.")
	fs.Var(&path_svg_alt, "path-svg-alt", "Zero or more alternate paths to serve SVG requests from.")

	fs.IntVar(&navplace_max_features, "navplace-max-features", 10, "The maximum number of WOF IDs allowed in a NavPlace request.")

	fs.BoolVar(&enable_cors, "enable-cors", false, "Enable CORS support.")
	fs.Var(&cors_allowed_origins, "cors-allowed-origin", "Zero or more allowed origins for CORS requests.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "A simple HTTP server-based interface for serving different machine-reabable representations (derivatives) of Who's On First documents.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s[options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	return fs
}
