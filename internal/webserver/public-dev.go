// +build dev

package webserver

import (
	"net/http"
)

var PublicFiles http.Handler = http.FileServer(Assets)
