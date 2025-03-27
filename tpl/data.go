// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	_ "chesslovaquia.github.io/go/clvq/cfg"
)

var _ Data = &BaseData{}

type Data interface{}

type BaseData struct{}

func NewData(path string) *BaseData {
	return &BaseData{}
}
