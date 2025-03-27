// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"embed"
)

//go:embed tpl
var fs embed.FS

type Tpl struct{}

func NewTpl() *Tpl {
	return &Tpl{}
}
