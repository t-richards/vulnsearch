package app

import (
	"encoding/json"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/gofiber/fiber/v2"
	"github.com/t-richards/vulnsearch/internal/assets"
	"github.com/t-richards/vulnsearch/internal/views"
)

func index(c *fiber.Ctx) error {
	return c.Render("views/index", nil, "views/layouts/main")
}

func build(c *fiber.Ctx) error {
	src, err := assets.Assets.ReadFile("public/application.ts")
	if err != nil {
		return err
	}
	opts := api.TransformOptions{
		LogLevel:          api.LogLevelInfo,
		Loader:            api.LoaderTS,
		Sourcemap:         api.SourceMapInline,
		Format:            api.FormatIIFE,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
	}

	result := api.Transform(string(src), opts)
	if len(result.Errors) > 0 {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSONP(result.Errors)
	}

	if len(result.Warnings) > 0 {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSONP(result.Warnings)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJavaScript)
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
	return c.Render("views/product", viewData, "views/layouts/main")
}
