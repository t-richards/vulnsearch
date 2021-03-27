package views

import (
	"embed"
	_ "embed" // Assets
	"fmt"
	"html/template"
	"io"
	"sync"
)

//go:embed templates/*.html
var templateSrc embed.FS

type EmbedEngine struct {
	// determines if the engine parsed all templates
	loaded bool
	// lock for funcmap and templates
	mutex sync.RWMutex
	// template functions
	funcmap map[string]interface{}
	// templates
	Templates *template.Template
}

func New() *EmbedEngine {
	engine := &EmbedEngine{}
	engine.AddFunc("layout", func() error {
		return fmt.Errorf("layout: called unexpectedly")
	})
	return engine
}

// AddFunc adds the function to the template's function map.
// It is legal to overwrite elements of the default actions
func (e *EmbedEngine) AddFunc(name string, fn interface{}) *EmbedEngine {
	e.mutex.Lock()
	e.funcmap[name] = fn
	e.mutex.Unlock()
	return e
}

func (e *EmbedEngine) Load() error {
	if e.loaded {
		return nil
	}

	var err error
	e.Templates, err = template.ParseFS(templateSrc, "templates/*.html")
	if err != nil {
		return err
	}

	e.loaded = true
	return nil
}

func (e *EmbedEngine) Render(out io.Writer, template string, binding interface{}, layout ...string) error {
	if !e.loaded {
		err := e.Load()
		if err != nil {
			return err
		}
	}

	tmpl := e.Templates.Lookup(template)
	if tmpl == nil {
		return fmt.Errorf("render: template %s does not exist", template)
	}

	if len(layout) > 0 {
		lay := e.Templates.Lookup(layout[0])
		if lay == nil {
			return fmt.Errorf("render: layout %s does not exist", layout[0])
		}
		lay.Funcs(map[string]interface{}{
			"layout": func() error {
				return tmpl.Execute(out, binding)
			},
		})
		return lay.Execute(out, binding)
	}
	return tmpl.Execute(out, binding)
}
