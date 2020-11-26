// +build dev

package webserver

import (
	"net/http"
)

var Assets = http.Dir("public")
var PublicFiles http.Handler = http.FileServer(Assets)
