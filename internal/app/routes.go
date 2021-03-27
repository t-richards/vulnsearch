package app

import (
	"encoding/json"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/gofiber/fiber/v2"
	"github.com/t-richards/vulnsearch/internal/assets"
	"github.com/t-richards/vulnsearch/internal/views"
)

func index(c *fiber.Ctx) error {
	return c.Render("index", nil, "layout")
}

func build(c *fiber.Ctx) error {
	src, err := assets.Assets.ReadFile("public/application.ts")
	if err != nil {
		return nil
	}
	opts := api.TransformOptions{
		LogLevel:          api.LogLevelInfo,
		Target:            api.ESNext,
		Loader:            api.LoaderTS,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Format:            api.FormatIIFE,
	}

	result := api.Transform(string(src), opts)
	if len(result.Errors) > 0 {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Send(result.Code)
}

func vendor(c *fiber.Ctx) error {
	params := SearchParams{}
	err := json.Unmarshal(c.Body(), &params)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := VendorResponse{
		Vendors: []string{},
	}
	conn.
		Table("products").
		Select("DISTINCT vendor").
		Where("vendor LIKE ?", *params.Vendor+"%").
		Order("vendor ASC").
		Limit(100).
		Find(&resp.Vendors)

	return c.JSON(resp)
}

func product(c *fiber.Ctx) error {
	params := SearchParams{}
	err := json.Unmarshal(c.Body(), &params)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := ProductResponse{
		Products: make([]string, 0),
	}
	conn.
		Table("products").
		Select("DISTINCT name").
		Where("vendor = ?", *params.Vendor).
		Where("name LIKE ?", *params.Name+"%").
		Order("name ASC").
		Limit(100).
		Find(&resp.Products)

	return c.JSON(resp)
}

func version(c *fiber.Ctx) error {
	params := SearchParams{}
	err := json.Unmarshal(c.Body(), &params)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := VersionResponse{
		Versions: make([]string, 0),
	}
	conn.
		Table("products").
		Select("DISTINCT version").
		Where("vendor = ?", *params.Vendor).
		Where("name = ?", *params.Name).
		Where("version LIKE ?", *params.Version+"%").
		Order("version ASC").
		Limit(100).
		Find(&resp.Versions)

	return c.JSON(resp)
}

func search(c *fiber.Ctx) error {
	viewData := views.ProductView{}
	// Load product
	conn.
		Where(
			"vendor = ? AND name = ? AND version = ?",
			c.Params("vendor"),
			c.Params("name"),
			c.Params("version"),
		).
		First(&viewData.Product)

	// Load CVEs
	conn.
		Model(&viewData.Product).
		Association("Cves").
		Find(&viewData.Cves)

	// Compute extra view data & render template
	viewData.Prepare()
	return c.Render("product", viewData, "layout")
}
