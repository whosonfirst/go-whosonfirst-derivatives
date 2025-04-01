package server

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-representation/http"
)

type RunOptions struct {
	ServerURI        string     `json:"server_uri"`
	SourceURI        string     `json:"source_uri"`
	AuthenticatorURI string     `json:"authenticator_uri"`
	URIs             *http.URIs `json:"uris"`
	Verbose          bool       `json:"verbose"`
}

func RunOptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "WHOSONFIRST")

	if err != nil {
		return nil, fmt.Errorf("Failed to assign flags from environment variables, %w", err)
	}

	uris := http.DefaultURIs()

	opts := &RunOptions{
		ServerURI:        server_uri,
		AuthenticatorURI: authenticator_uri,
		SourceURI:        source_uri,
		URIs:             uris,
		Verbose:          verbose,
	}

	return opts, nil
}
