// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"embed"
	"html/template"
	"path/filepath"
	"strings"
	"sync"

	"chesslovaquia.github.io/go/clvq/tpl"

	_ "chesslovaquia.github.io/go/clvq/cfg"
)

var _ tpl.Tpl = &Tpl{}
var _ tpl.Data = &TplData{}

//go:embed static
var StaticFS embed.FS

//go:embed tpl
var fs embed.FS

type TplData struct {
	Root string
	Site string
}

func newTplData(path string) *TplData {
	return &TplData{
		Root: "/_",
		Site: "clvq",
	}
}

type Tpl struct {
	mutex sync.Mutex
}

func NewTpl() *Tpl {
	return &Tpl{}
}

func (t *Tpl) BaseFile() string {
	return filepath.Join("tpl", "base.html")
}

func (t *Tpl) Filepath(path string) string {
	return filepath.Join("tpl", strings.TrimPrefix(path, "/_/"))
}

func (t *Tpl) Get(path string) (*template.Template, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	tmplt, err := template.ParseFS(fs, t.BaseFile(), t.Filepath(path))
	if err != nil {
		return nil, err
	}
	return tmplt, nil
}

func (t *Tpl) GetData(path string) tpl.Data {
	return newTplData(path)
}
