// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tpl

import (
	"html/template"
	"path/filepath"
	"sync"

	"chesslovaquia.github.io/go/clvq/cfg"
)

var _ Tpl = &TplBase{}

type Tpl interface {
	BaseFile() string
	Filepath(path string) string
	Get(path string) (*template.Template, error)
	GetData(path string) Data
}

type TplBase struct {
	mutex sync.Mutex
}

func New() *TplBase {
	return &TplBase{}
}

func (t *TplBase) BaseFile() string {
	return filepath.Join(cfg.TplDir(), cfg.TplBase())
}

func (t *TplBase) Filepath(path string) string {
	return filepath.Join(cfg.TplDir(), path)
}

func (t *TplBase) Get(path string) (*template.Template, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	tmplt, err := template.ParseFiles(t.BaseFile(), t.Filepath(path))
	if err != nil {
		return nil, err
	}
	return tmplt, nil
}

func (t *TplBase) GetData(path string) Data {
	return NewData(path)
}
