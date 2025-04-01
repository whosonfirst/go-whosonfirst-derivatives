package http

import (
	"context"
	"fmt"
	"io"
	go_http "net/http"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/aaronland/go-http-sanitize"
	"github.com/whosonfirst/go-whosonfirst-representation"
	wof_uri "github.com/whosonfirst/go-whosonfirst-uri"
)

var re_path_id = regexp.MustCompile(`/id/(\d+)/.*$`)

type URI struct {
	Id          int64
	URI         string
	Feature     []byte
	URIArgs     *wof_uri.URIArgs
	IsAlternate bool
}

func ParseURIFromRequest(req *go_http.Request) (*URI, error, int) {

	ctx := req.Context()

	path, err := sanitize.GetString(req, "id")

	if err != nil {
		return nil, fmt.Errorf("Failed to derive ?id= parameter, %w", err), go_http.StatusBadRequest
	}

	/*

		resource, err := sanitize.GetString(req, "resource")

		if err != nil {
			return nil, fmt.Errorf("Failed to derive ?resource= parameter, %w", err), go_http.StatusBadRequest
		}

		if path == "" && resource != "" {

			wof_uri, err := webfinger.DeriveWhosOnFirstURIFromResource(resource)

			if err != nil {

			}

			path = wof_uri
		}
	*/

	// Y U NO WORK...
	// https://pkg.go.dev/net/http@master#hdr-Patterns

	if path == "" {
		path = req.PathValue("id")
	}

	// Oh well, at least the ServeMux recognizes wildcards now...
	if path == "" {

		path = req.URL.Path

		if re_path_id.MatchString(path) {
			m := re_path_id.FindStringSubmatch(path)
			path = m[1]
		}
	}

	return ParseURIFromPath(ctx, path)
}

func ParseURIFromPath(ctx context.Context, path string) (*URI, error, int) {

	wofid, uri_args, err := wof_uri.ParseURI(path)

	if err != nil {
		return nil, fmt.Errorf("Error locating record for %s, %w", path, err), go_http.StatusNotFound
	}

	if wofid == -1 {
		return nil, fmt.Errorf("Failed to locate record for %s", path), go_http.StatusNotFound
	}

	fname, err := wof_uri.Id2Fname(wofid, uri_args)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive filename from %d (%s), %w", wofid, path, err), go_http.StatusInternalServerError
	}

	ext := filepath.Ext(fname)
	fname = strings.Replace(fname, ext, "", 1)

	parsed_uri := &URI{
		Id:          wofid,
		URI:         fname,
		URIArgs:     uri_args,
		IsAlternate: uri_args.IsAlternate,
	}

	return parsed_uri, nil, 0
}

func FeatureFromRequestURI(ctx context.Context, src representation.Source, req_uri *URI) ([]byte, error) {

	wof_id := req_uri.Id

	r, err := src.GetFeature(ctx, wof_id, req_uri.URIArgs)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve feature for %d, %w", wof_id, err)
	}

	defer r.Close()

	return io.ReadAll(r)
}
