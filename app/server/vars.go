package server

import (
	"sync"

	"github.com/rs/cors"
	"github.com/sfomuseum/go-http-auth"
	"github.com/whosonfirst/go-whosonfirst-representation"
	"github.com/whosonfirst/go-whosonfirst-representation/http"
)

var run_options *RunOptions

var src representation.Source

var authenticator auth.Authenticator

var uris_table *http.URIs

var setupCommonOnce sync.Once
var setupCommonError error

var setupAPIOnce sync.Once
var setupAPIError error

var cors_wrapper *cors.Cors
