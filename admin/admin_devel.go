// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"path/filepath"
	"strings"

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
	*tpl.TplBase
}

func newTplDev(path string) *TplDev {
	return &TplDev{tpl.New()}
}

func (t *TplDev) BaseFile() string {
	return filepath.Join("admin", "tpl", "base.html")
}

func (t *TplDev) Filepath(path string) string {
	return filepath.Join("admin", "tpl", strings.TrimPrefix(path, "/_/"))
}

func (t *TplDev) GetData(path string) tpl.Data {
	return newTplDevData(path)
}
