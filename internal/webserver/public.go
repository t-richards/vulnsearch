// +build !dev

package webserver

import (
	"net/http"

	"github.com/shurcooL/httpgzip"
)

var publicFiles http.Handler = httpgzip.FileServer(Assets, httpgzip.FileServerOptions{})
