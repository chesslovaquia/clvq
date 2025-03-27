// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"html/template"
	"path/filepath"

	"chesslovaquia.github.io/go/clvq/cfg"
)

var _ Tpl = &TplBase{}

type Tpl interface {
	BaseFile() string
	Filepath(path string) string
	Get(path string) (*template.Template, error)
	GetData(path string) Data
}

type TplBase struct{}

func New() *TplBase {
	return &TplBase{}
}

func (t *TplBase) BaseFile() string {
	return "FIXME.html"
}

func (t *TplBase) Filepath(path string) string {
	return filepath.Join(cfg.Tpl.Dir, path)
}

func (t *TplBase) Get(path string) (*template.Template, error) {
	return nil, nil
}

func (t *TplBase) GetData(path string) Data {
	return NewData(path)
}
