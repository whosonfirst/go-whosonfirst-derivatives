package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/rs/cors"
	"github.com/sfomuseum/go-http-auth"
	"github.com/whosonfirst/go-whosonfirst-representation"
)

func setupCommon() {

	ctx := context.Background()
	var err error

	// defined in vars.go
	src, err = representation.NewSource(ctx, run_options.SourceURI)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to set up source, %w", err)
		return
	}

	authenticator, err = auth.NewAuthenticator(ctx, run_options.AuthenticatorURI)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to set up authenticator, %w", err)
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

	// Please finesse me...
	cors_origins := []string{
		"*",
	}

	cors_wrapper = cors.New(cors.Options{
		AllowedOrigins:   cors_origins,
		AllowCredentials: false,
		Debug:            false,
	})
}
