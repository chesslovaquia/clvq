// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package site

import (
	"embed"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

//go:embed static
var staticFS embed.FS

func ServeStaticFS(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(path.Clean(r.URL.Path), "/clvq/")
	fn := filepath.Join("static", p)
	log.Printf("000 %s - site:%s", p, fn)
	http.ServeFileFS(w, r, staticFS, fn)
}
