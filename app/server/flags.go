package server

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var server_uri string
var source_uri string
var authenticator_uri string

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("server")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "...")
	fs.StringVar(&source_uri, "source-uri", "reader://?reader-uri=https://data.whosonfirst.org", "...")
	fs.StringVar(&authenticator_uri, "authenticator-uri", "null://", "...")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
