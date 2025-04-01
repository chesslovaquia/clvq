// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package site

import (
	"chesslovaquia.github.io/go/clvq"
	"chesslovaquia.github.io/go/clvq/tpl"
)

func Main() {
	clvq.AddHandler("/", tpl.New())
	clvq.Main()
}
