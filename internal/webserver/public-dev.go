// +build dev

package webserver

import (
	"net/http"
)

var publicFiles http.Handler = http.FileServer(Assets)
