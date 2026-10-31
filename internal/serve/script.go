package serve

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/evanw/esbuild/pkg/api"

	"github.com/t-richards/vulnsearch/assets"
)

// errScriptBuild is returned when the TypeScript source fails to compile.
var errScriptBuild = errors.New("failed to build application.js")

// script is compiled JavaScript, ready to serve.
type script struct {
	code []byte
	etag string
}

// buildScript compiles the embedded TypeScript source. The source is fixed at
// build time, so this only needs to run once per process.
func buildScript() (script, error) {
	result := api.Transform(assets.ApplicationTS, api.TransformOptions{
		Sourcefile:        "application.ts",
		Loader:            api.LoaderTS,
		Format:            api.FormatESModule,
		Target:            api.ES2022,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
	})

	if len(result.Errors) > 0 {
		msg := result.Errors[0]

		return script{}, fmt.Errorf("%w: %s", errScriptBuild, msg.Text)
	}

	sum := sha256.Sum256(result.Code)

	return script{
		code: result.Code,
		etag: `"` + hex.EncodeToString(sum[:8]) + `"`,
	}, nil
}

// serveScript serves the compiled script. Browsers revalidate with the ETag,
// so an upgraded binary's script is picked up immediately.
func (app *Application) serveScript(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("ETag", app.script.etag)
	http.ServeContent(w, r, "application.js", time.Time{}, bytes.NewReader(app.script.code))
}
