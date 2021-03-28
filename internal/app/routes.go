package app

import (
	"encoding/json"
	"os"

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
		Format:            api.FormatIIFE,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
	}

	_, ok := os.LookupEnv("DEBUG")
	if ok {
		opts.Sourcemap = api.SourceMapInline
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJavaScript)
	result := api.Transform(string(src), opts)
	if len(result.Errors) > 0 {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(result.Warnings) > 0 {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Send(result.Code)
}

func vendor(c *fiber.Ctx) error {
	params := SearchParams{}
	err := json.Unmarshal(c.Body(), &params)
	if err != nil {
		return err
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
		return err
	}

	resp := ProductResponse{
		Products: []string{},
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
		return err
	}

	resp := VersionResponse{
		Versions: []string{},
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
	productResult :=
		conn.
			Where(
				"vendor = ? AND name = ? AND version = ?",
				c.Query("vendor"),
				c.Query("name"),
				c.Query("version"),
			).
			First(&viewData.Product)

	if productResult.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Load CVEs
	conn.
		Model(&viewData.Product).
		Association("Cves").
		Find(&viewData.Cves)

	// Compute extra view data & render template
	viewData.Prepare()
	return c.Render("views/product", viewData, "views/layouts/main")
}
