package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"

	"github.com/whosonfirst/go-whosonfirst-derivatives/http/api"
)

func geoJSONHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupAPIOnce.Do(setupAPI)

	if setupAPIError != nil {
		slog.Error("Failed to set up common configuration", "error", setupAPIError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupAPIError)
	}

	opts := &api.GeoJSONHandlerOptions{
		Provider: prv,
	}

	h, err := api.GeoJSONHandler(opts)

	if err != nil {
		return nil, err
	}

	return cors_wrapper.Handler(h), nil
}

func geoJSONLDHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupAPIOnce.Do(setupAPI)

	if setupAPIError != nil {
		slog.Error("Failed to set up common configuration", "error", setupAPIError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupAPIError)
	}

	opts := &api.GeoJSONLDHandlerOptions{
		Provider: prv,
	}

	h, err := api.GeoJSONLDHandler(opts)

	if err != nil {
		return nil, err
	}

	return cors_wrapper.Handler(h), nil
}

func sprHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupAPIOnce.Do(setupAPI)

	if setupAPIError != nil {
		slog.Error("Failed to set up common configuration", "error", setupAPIError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupAPIError)
	}

	opts := &api.SPRHandlerOptions{
		Provider: prv,
	}

	h, err := api.SPRHandler(opts)

	if err != nil {
		return nil, err
	}

	return cors_wrapper.Handler(h), nil
}

func selectHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupAPIOnce.Do(setupAPI)

	if setupAPIError != nil {
		slog.Error("Failed to set up common configuration", "error", setupAPIError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupAPIError)
	}

	// Make this a config/flag
	select_pattern := `properties(?:.[a-zA-Z0-9-_]+){1,}`

	pat, err := regexp.Compile(select_pattern)

	if err != nil {
		slog.Error("Failed to compile select pattern", "pattern", select_pattern, "error", err)
		return nil, fmt.Errorf("Failed to compile select pattern (%s), %w", select_pattern, err)
	}

	opts := &api.SelectHandlerOptions{
		Pattern:  pat,
		Provider: prv,
	}

	h, err := api.SelectHandler(opts)

	if err != nil {
		return nil, err
	}

	return cors_wrapper.Handler(h), nil
}

func navPlaceHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupAPIOnce.Do(setupAPI)

	if setupAPIError != nil {
		slog.Error("Failed to set up common configuration", "error", setupAPIError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupAPIError)
	}

	opts := &api.NavPlaceHandlerOptions{
		Provider:    prv,
		MaxFeatures: 10,
	}

	h, err := api.NavPlaceHandler(opts)

	if err != nil {
		return nil, err
	}

	return cors_wrapper.Handler(h), nil
}

func svgHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupAPIOnce.Do(setupAPI)

	if setupAPIError != nil {
		slog.Error("Failed to set up common configuration", "error", setupAPIError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupAPIError)
	}

	sz := api.DefaultSVGSizes()

	opts := &api.SVGHandlerOptions{
		Provider: prv,
		Sizes:    sz,
	}

	return api.SVGHandler(opts)
}
