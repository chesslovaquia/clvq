// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"chesslovaquia.github.io/go/clvq"
	"chesslovaquia.github.io/go/clvq/tpl"
)

func main() {
	clvq.AddHandler("/", tpl.New())
	clvq.Main()
}
