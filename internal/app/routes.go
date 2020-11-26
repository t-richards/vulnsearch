package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/t-richards/vulnsearch/internal/views"
)

func (app *App) index() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := views.HomeTemplate.Execute(w, nil)
		if err != nil {
			log.Printf("Error rendering home template: %v", err)
		}
	}
}

// TODO(tom): DRY this up
func (app *App) vendor() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		params := SearchParams{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer r.Body.Close()

		resp := VendorResponse{
			Vendors: make([]string, 0),
		}
		app.DB.
			Table("products").
			Select("DISTINCT vendor").
			Where("vendor LIKE ?", *params.Vendor+"%").
			Order("vendor ASC").
			Limit(100).
			Find(&resp.Vendors)

		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding vendor response: %v", err)
		}
	}
}

func (app *App) product() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		params := SearchParams{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer r.Body.Close()

		resp := ProductResponse{
			Products: make([]string, 0),
		}
		app.DB.
			Table("products").
			Select("DISTINCT name").
			Where("vendor = ?", *params.Vendor).
			Where("name LIKE ?", *params.Name+"%").
			Order("name ASC").
			Limit(100).
			Find(&resp.Products)

		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding vendor response: %v", err)
		}
	}
}

func (app *App) version() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		params := SearchParams{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer r.Body.Close()

		resp := VersionResponse{
			Versions: make([]string, 0),
		}
		app.DB.
			Table("products").
			Select("DISTINCT version").
			Where("vendor = ?", *params.Vendor).
			Where("name = ?", *params.Name).
			Where("version LIKE ?", *params.Version+"%").
			Order("version ASC").
			Limit(100).
			Find(&resp.Versions)

		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding vendor response: %v", err)
		}
	}
}
