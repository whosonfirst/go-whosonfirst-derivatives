package api

import (
	"net/http"

	"github.com/sfomuseum/go-geojsonld"
	"github.com/whosonfirst/go-whosonfirst-derivatives"
	wof_http "github.com/whosonfirst/go-whosonfirst-derivatives/http"
)

type GeoJSONLDHandlerOptions struct {
	Provider derivatives.Provider
}

func GeoJSONLDHandler(opts *GeoJSONLDHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := wof_http.LoggerWithRequest(req, nil)

		req_uri, err, status := wof_http.ParseURIFromRequest(req)

		if err != nil {
			logger.Error("Failed to parse URI from request", "error", err)
			http.Error(rsp, derivatives.ErrNotFound.Error(), status)
			return
		}

		if req_uri.Id <= -1 {
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		logger = logger.With("id", req_uri.Id)

		r, err := wof_http.FeatureFromRequestURI(ctx, opts.Provider, req_uri)

		if err != nil {
			logger.Error("Failed to get by ID", "error", err)
			http.Error(rsp, derivatives.ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		body, err := geojsonld.AsGeoJSONLD(ctx, r)

		if err != nil {
			logger.Error("Failed to render geojson", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/geo+json")
		rsp.Write([]byte(body))
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
