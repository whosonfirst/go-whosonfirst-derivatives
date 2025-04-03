package server

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"strings"

	"github.com/mitchellh/copystructure"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-derivatives/http"
)

type RunOptions struct {
	ServerURI           string     `json:"server_uri"`
	ProviderURI         string     `json:"provider_uri"`
	URIs                *http.URIs `json:"uris"`
	NavPlaceMaxFeatures int        `json:"navplace_max_features"`
	EnableCORS          bool       `json:"enable_cors"`
	CORSAllowedOrigins  []string   `json:"cors_allowed_origins"`
	Verbose             bool       `json:"verbose"`
}

func (o *RunOptions) Clone() (*RunOptions, error) {

	v, err := copystructure.Copy(o)

	if err != nil {
		return nil, fmt.Errorf("Failed to create local run options, %w", err)
	}

	new_opts := v.(*RunOptions)
	return new_opts, nil
}

func RunOptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "WHOSONFIRST")

	if err != nil {
		return nil, fmt.Errorf("Failed to assign flags from environment variables, %w", err)
	}

	uris := http.DefaultURIs()
	uris.GeoJSON = path_geojson
	uris.GeoJSONAlt = path_geojson_alt
	uris.GeoJSONLD = path_geojsonld
	uris.GeoJSONLDAlt = path_geojsonld_alt
	uris.NavPlace = path_navplace
	uris.NavPlaceAlt = path_navplace_alt
	uris.Select = path_select
	uris.SelectAlt = path_select_alt
	uris.SPR = path_spr
	uris.SPRAlt = path_spr_alt
	uris.SVG = path_svg
	uris.SVGAlt = path_svg_alt
	uris.WKT = path_wkt
	uris.WKTAlt = path_wkt_alt

	if reader_uri != "" && strings.Contains(provider_uri, "{reader_uri}") {
		provider_uri = strings.Replace(provider_uri, "{reader_uri}", url.QueryEscape(reader_uri), 1)
	}

	opts := &RunOptions{
		ServerURI:           server_uri,
		ProviderURI:         provider_uri,
		URIs:                uris,
		NavPlaceMaxFeatures: navplace_max_features,
		EnableCORS:          enable_cors,
		CORSAllowedOrigins:  cors_allowed_origins,
		Verbose:             verbose,
	}

	return opts, nil
}
