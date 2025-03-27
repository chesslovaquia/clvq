// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"embed"
	"html/template"
	"path/filepath"
	"sync"

	"chesslovaquia.github.io/go/clvq/tpl"

	_ "chesslovaquia.github.io/go/clvq/cfg"
)

var _ tpl.Tpl = &Tpl{}

//go:embed static
var StaticFS embed.FS

//go:embed tpl
var fs embed.FS

type Tpl struct {
	mutex sync.Mutex
	path  string
}

func NewTpl() *Tpl {
	return &Tpl{
		path: "/_.html",
	}
}

func (t *Tpl) BaseFile() string {
	return filepath.Join("tpl", "base.html")
}

func (t *Tpl) Path() string {
	return filepath.Join("tpl", t.path)
}

func (t *Tpl) load(path string) (*template.Template, error) {
	t.path = path
	tmplt, err := template.ParseFS(fs, t.BaseFile(), t.Path())
	if err != nil {
		return nil, err
	}
	return tmplt, nil
}

func (t *Tpl) Get(path string) (*template.Template, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.load(path)
}
