// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package admin

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"chesslovaquia.github.io/go/clvq/cfg"
	"chesslovaquia.github.io/go/clvq/tpl"
)

var _ tpl.Tpl = &Tpl{}
var _ tpl.Data = &TplData{}

//go:embed static
var StaticFS embed.FS

//go:embed tpl
var fs embed.FS

// TplData

type TplData struct{}

func newTplData(path string) *TplData {
	return &TplData{}
}

func (d *TplData) Root() string {
	return cfg.Tpl.Root
}

func (d *TplData) Site() string {
	return cfg.Tpl.Site
}

// Tpl

type Tpl struct {
	mutex sync.Mutex
}

func NewTpl() *Tpl {
	return &Tpl{}
}

func (t *Tpl) BaseFile() string {
	return filepath.Join("tpl", "base.html")
}

func (t *Tpl) Filepath(path string) string {
	return filepath.Join("tpl", strings.TrimPrefix(path, "/_/"))
}

func (t *Tpl) Get(path string) (*template.Template, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	tmplt, err := template.ParseFS(fs, t.BaseFile(), t.Filepath(path))
	if err != nil {
		return nil, err
	}
	return tmplt, nil
}

func (t *Tpl) GetData(path string) tpl.Data {
	return newTplData(path)
}

// config.json

func ConfigJSONHandler(w http.ResponseWriter, r *http.Request) {
	mimeType := mime.TypeByExtension(".json")
	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cfg.Get()); err != nil {
		log.Printf("500 /_/config.json - %v", err)
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("200 /_/config.json")
}
