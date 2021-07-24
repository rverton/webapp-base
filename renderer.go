package main

import (
	"embed"
	"errors"
	"io"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

//go:embed templates/*
var content embed.FS

type Renderer struct {
	templateSet *pongo2.TemplateSet
	Debug       bool
}

type PongoLoader struct {
	virtualFs embed.FS
}

// Abs calculates the path to a given template. Whenever a path must be resolved
// due to an import from another template, the base equals the parent template's path.
func (p PongoLoader) Abs(base, name string) string {
	absPath := name

	if base != "" {
		absPath = filepath.Join(filepath.Dir(base), name)
	}

	return absPath
}

func (p PongoLoader) Get(path string) (io.Reader, error) {
	return p.virtualFs.Open(path)
}

func NewRenderer(debug bool) *Renderer {
	pl := &PongoLoader{
		virtualFs: content,
	}
	return &Renderer{
		Debug:       debug,
		templateSet: pongo2.NewSet("assimilator", pl),
	}
}

func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var ctx pongo2.Context
	var err error
	var t *pongo2.Template

	if data != nil {
		var ok bool
		if ctx, ok = data.(pongo2.Context); !ok {
			return errors.New("no pongo2.Context data was passed")
		}
	}

	// if in debug mode, load file from fs, otherwise load from cache/virtual fs
	if r.Debug {
		t, err = pongo2.FromFile(name)
	} else {
		t, err = r.templateSet.FromCache(name)
	}

	if err != nil {
		return err
	}
	return t.ExecuteWriter(ctx, w)
}
