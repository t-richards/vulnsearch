package serve

import (
	"net/http"
	"strings"
)

// contentSecurityPolicy allows the app's own scripts, plus Bootstrap's CSS
// from the CDN. Bootstrap embeds a few SVG icons as data: URIs.
var contentSecurityPolicy = strings.Join([]string{
	"default-src 'none'",
	"base-uri 'none'",
	"frame-ancestors 'none'",
	"connect-src 'self'",
	"form-action 'self'",
	"img-src 'self' data:",
	"script-src 'self'",
	"style-src https://cdn.jsdelivr.net",
}, "; ")

// securityHeaders sets HTTP headers for security purposes.
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", contentSecurityPolicy)
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}
