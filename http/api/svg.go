package api

import (
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-sanitize"
	"github.com/whosonfirst/go-whosonfirst-derivatives"
	wof_http "github.com/whosonfirst/go-whosonfirst-derivatives/http"
	"github.com/whosonfirst/go-whosonfirst-svg"
)

type SVGSize struct {
	Label     string
	MaxHeight int
	MaxWidth  int
}

type SVGHandlerOptions struct {
	Sizes    map[string]SVGSize
	Provider derivatives.Provider
}

func DefaultSVGSizes() map[string]SVGSize {

	sm := SVGSize{
		Label:     "sm",
		MaxHeight: 300,
		MaxWidth:  300,
	}

	med := SVGSize{
		Label:     "med",
		MaxHeight: 640,
		MaxWidth:  640,
	}

	lg := SVGSize{
		Label:     "lg",
		MaxHeight: 1024,
		MaxWidth:  1024,
	}

	sz := map[string]SVGSize{
		"sm":  sm,
		"med": med,
		"lg":  lg,
	}

	return sz
}

func SVGHandler(opts *SVGHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := wof_http.LoggerWithRequest(req, nil)

		req_uri, err, status := wof_http.ParseURIFromRequest(req)

		if err != nil {

			logger.Error("Failed to parse URI from request", "error", err)

			http.Error(rsp, err.Error(), status)
			return
		}

		if req_uri.Id <= -1 {
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		logger = logger.With("id", req_uri.Id)

		f, err := wof_http.FeatureFromRequestURI(ctx, opts.Provider, req_uri)

		if err != nil {
			logger.Error("Failed to get by ID", "id", req_uri.Id, "error", err)
			http.Error(rsp, derivatives.ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		sn_opts := sanitize.DefaultOptions()

		sz := "lg"

		query := req.URL.Query()
		query_sz := query.Get("size")

		req_sz, err := sanitize.SanitizeString(query_sz, sn_opts)

		if err != nil {
			http.Error(rsp, err.Error(), status)
			return
		}

		if req_sz != "" {
			sz = req_sz
		}

		sz_info, ok := opts.Sizes[sz]

		if !ok {
			http.Error(rsp, "Invalid output size", http.StatusBadRequest)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Header().Set("Content-Type", "image/svg+xml")

		opts := svg.NewDefaultOptions()
		opts.Height = float64(sz_info.MaxHeight)
		opts.Width = float64(sz_info.MaxWidth)
		opts.Writer = rsp

		// to do: support for custom styles:
		// https://github.com/whosonfirst/go-whosonfirst-browser/v6issues/19

		opts.StyleFunction = func(f []byte) (map[string]string, error) {

			attrs := make(map[string]string)

			type_rsp := gjson.GetBytes(f, "geometry.type")

			if !type_rsp.Exists() {
				return nil, fmt.Errorf("Missing geometry.type")
			}

			geom_type := type_rsp.String()
			// log.Println(geom_type)

			switch geom_type {
			case "LineString":
				attrs["fill-opacity"] = "0.0"
				attrs["stroke-width"] = "1.0"
				attrs["stroke-opacity"] = "2.0"
				attrs["stroke"] = "#000"
			case "Point", "MultiPoint":
				// something something something
				// https://github.com/whosonfirst/go-whosonfirst-browser/v6issues/18
			default:
				// pass
			}

			return attrs, nil
		}

		svg.FeatureToSVG(f, opts)
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
