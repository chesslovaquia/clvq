// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"chesslovaquia.github.io/go/clvq/cfg"
)

var _ Data = &BaseData{}

type Data interface {
	Root() string
	Site() string
}

type BaseData struct{}

func NewData(path string) *BaseData {
	return &BaseData{}
}

func (d *BaseData) Root() string {
	return cfg.Tpl.Root
}

func (d *BaseData) Site() string {
	return cfg.Tpl.Site
}
