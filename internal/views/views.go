package views

import (
	"html/template"
)

// HomeTemplate is the template for the homepage
var HomeTemplate *template.Template

// ProductTemplate is the template for the product page
var ProductTemplate *template.Template

func init() {
	HomeTemplate = parseTemplateWithLayout("home", homeSrc)
	ProductTemplate = parseTemplateWithLayout("product", productSrc)
}

func parseTemplateWithLayout(tmplName string, tmplSrc string) *template.Template {
	tpl := template.Must(template.New("layout").Parse(layoutSrc))
	template.Must(tpl.New(tmplName).Parse(tmplSrc))

	return tpl
}
