// +build !dev

package webserver

import (
	"net/http"

	"github.com/shurcooL/httpgzip"
)

// PublicFiles serves pre-gzipped asset content from the generated assets-prod.go file
var PublicFiles http.Handler = httpgzip.FileServer(Assets, httpgzip.FileServerOptions{})
