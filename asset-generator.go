// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	fs := http.Dir("public")
	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "internal/webserver/assets-prod.go",
		PackageName:  "webserver",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
