package server

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var server_uri string
var provider_uri string
var authenticator_uri string

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("server")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A registered aaronland/go-http-server.Server URI.")
	fs.StringVar(&provider_uri, "provider-uri", "reader://?reader-uri=https://data.whosonfirst.org", "A registered whosonfirst/go-whosonfirst-derivatives.Provider URI.")
	fs.StringVar(&authenticator_uri, "authenticator-uri", "null://", "A registered sfomuseum/go-auth.Authenticator URI.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
