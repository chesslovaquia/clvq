// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"html/template"

	_ "chesslovaquia.github.io/go/clvq/cfg"
)

var _ Tpl = &TplBase{}

type Tpl interface {
	Get(path string) (*template.Template, error)
}

type TplBase struct{}

func New() *TplBase {
	return &TplBase{}
}

func (t *TplBase) Get(path string) (*template.Template, error) {
	return nil, nil
}
