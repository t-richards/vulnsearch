// +build !dev

package webserver

import (
	"net/http"

	"github.com/shurcooL/httpgzip"
)

var PublicFiles http.Handler = httpgzip.FileServer(Assets, httpgzip.FileServerOptions{})
