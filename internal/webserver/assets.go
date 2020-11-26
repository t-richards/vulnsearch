// +build dev

package webserver

import (
	"net/http"
)

var Assets = http.Dir("public/dist")
