package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-server/handler"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to derive run options from flagset, %w", err)
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	// START OF defer loading handlers (and all their dependencies) until they are actually routed to
	// in case we are running in a "serverless" environment like AWS Lambda

	handlers := map[string]handler.RouteHandlerFunc{

		// Common handler things
		// "/robots.txt": robotsTxtHandlerFunc,

		// API/machine-readable
		run_options.URIs.GeoJSON:   geoJSONHandlerFunc,
		run_options.URIs.GeoJSONLD: geoJSONLDHandlerFunc,
		run_options.URIs.NavPlace:  navPlaceHandlerFunc,
		run_options.URIs.Select:    selectHandlerFunc,
		run_options.URIs.SPR:       sprHandlerFunc,
		run_options.URIs.SVG:       svgHandlerFunc,
	}

	assign_handlers := func(handler_map map[string]handler.RouteHandlerFunc, paths []string, handler_func handler.RouteHandlerFunc) {

		for _, p := range paths {
			handler_map[p] = handler_func
		}
	}

	// API/machine-readable
	assign_handlers(handlers, run_options.URIs.GeoJSONAlt, geoJSONHandlerFunc)
	assign_handlers(handlers, run_options.URIs.GeoJSONLDAlt, geoJSONLDHandlerFunc)
	assign_handlers(handlers, run_options.URIs.NavPlaceAlt, navPlaceHandlerFunc)
	assign_handlers(handlers, run_options.URIs.SelectAlt, selectHandlerFunc)
	assign_handlers(handlers, run_options.URIs.SPRAlt, sprHandlerFunc)
	assign_handlers(handlers, run_options.URIs.SVGAlt, svgHandlerFunc)

	logger := slog.Default()
	
	log_logger := slog.NewLogLogger(logger.Handler(), slog.LevelInfo)

	route_handler_opts := &handler.RouteHandlerOptions{
		Handlers: handlers,
		Logger:   log_logger,
	}

	route_handler, err := handler.RouteHandlerWithOptions(route_handler_opts)

	if err != nil {
		return fmt.Errorf("Failed to configure route handler, %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", route_handler)

	// END OF defer loading handlers (and all their dependencies) until they are actually routed to

	s, err := server.NewServer(ctx, run_options.ServerURI)

	if err != nil {
		return fmt.Errorf("Failed to create new server, %w", err)
	}

	go func() {
		for uri, h := range handlers {
			slog.Debug("Enable handler", "uri", uri, "handler", fmt.Sprintf("%T", h))
		}
	}()

	slog.Info("Listening for requests", "address", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to start server, %w", err)
	}

	return nil
}
