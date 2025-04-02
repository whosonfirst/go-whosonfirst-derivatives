package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/rs/cors"
	"github.com/whosonfirst/go-whosonfirst-derivatives"
)

func setupCommon() {

	ctx := context.Background()
	var err error

	// defined in vars.go
	prv, err = derivatives.NewProvider(ctx, run_options.ProviderURI)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to set up provider, %w", err)
		return
	}
}

func setupAPI() {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		slog.Error("Failed to set up common configuration", "error", setupCommonError)
		setupAPIError = fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
		return
	}

	if run_options.EnableCORS {

		cors_wrapper = cors.New(cors.Options{
			AllowedOrigins:   run_options.CORSAllowedOrigins,
			AllowCredentials: false,
			Debug:            false,
		})
	}
}
