package webserver

import (
	"log"
	"net/http"
	"vulnsearch/internal/views"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	views.HomeTemplate.Execute(w, nil)
}

// Serve creates and starts the HTTP server
func Serve() {
	router := httprouter.New()
	router.GET("/", index)
	router.NotFound = publicFiles

	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
