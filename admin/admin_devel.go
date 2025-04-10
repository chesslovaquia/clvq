// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"path/filepath"
	"strings"
	"sync"
	"html/template"

	"chesslovaquia.github.io/go/clvq/cfg"
	"chesslovaquia.github.io/go/clvq/tpl"
)

var _ tpl.Tpl = &TplDev{}
var _ tpl.Data = &TplDevData{}

// TplDevData

type TplDevData struct{}

func newTplDevData(path string) *TplDevData {
	return &TplDevData{}
}

func (d *TplDevData) Root() string {
	return cfg.Tpl.Root
}

func (d *TplDevData) Site() string {
	return cfg.Tpl.Site
}

// TplDev

type TplDev struct {
	mutex sync.Mutex
}

func newTplDev() *TplDev {
	return &TplDev{}
}

func (t *TplDev) BaseFile() string {
	return filepath.Join("admin", "tpl", "base.html")
}

func (t *TplDev) Filepath(path string) string {
	return filepath.Join("admin", "tpl", strings.TrimPrefix(path, "/_/"))
}

func (t *TplDev) Get(path string) (*template.Template, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	tmplt, err := template.ParseFiles(t.BaseFile(), t.Filepath(path))
	if err != nil {
		return nil, err
	}
	return tmplt, nil
}

func (t *TplDev) GetData(path string) tpl.Data {
	return newTplDevData(path)
}
