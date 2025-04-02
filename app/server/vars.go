package server

import (
	"sync"

	"github.com/rs/cors"
	"github.com/whosonfirst/go-whosonfirst-derivatives"
	"github.com/whosonfirst/go-whosonfirst-derivatives/http"
)

var run_options *RunOptions

var prv derivatives.Provider

var uris_table *http.URIs

var setupCommonOnce sync.Once
var setupCommonError error

var setupAPIOnce sync.Once
var setupAPIError error

var cors_wrapper *cors.Cors
