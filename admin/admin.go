// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	_ "chesslovaquia.github.io/go/clvq/cfg"
)

type Tpl struct{}

func NewTpl() *Tpl {
	return &Tpl{}
}
