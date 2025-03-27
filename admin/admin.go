// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"embed"
	"html/template"

	"chesslovaquia.github.io/go/clvq/tpl"

	_ "chesslovaquia.github.io/go/clvq/cfg"
)

var _ tpl.Tpl = &Tpl{}

//go:embed static
var StaticFS embed.FS

//go:embed tpl
var fs embed.FS

type Tpl struct{}

func NewTpl() *Tpl {
	return &Tpl{}
}

func (t *Tpl) Get(path string) (*template.Template, error) {
	return nil, nil
}
