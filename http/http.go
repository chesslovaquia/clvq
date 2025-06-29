// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package http

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"chesslovaquia.github.io/go/clvq/admin"
	"chesslovaquia.github.io/go/clvq/cfg"
	"chesslovaquia.github.io/go/clvq/env"
	"chesslovaquia.github.io/go/clvq/site"
	"chesslovaquia.github.io/go/clvq/tpl"
)

func AddHandler(path string, template tpl.Tpl) {
	h := newHandler(template)
	http.HandleFunc(path, h.Handle)
}

type Handler struct {
	tpl tpl.Tpl
}

func newHandler(template tpl.Tpl) *Handler {
	return &Handler{
		tpl: template,
	}
}

func (h *Handler) ServeFile(w http.ResponseWriter, r *http.Request, path string) {
	var fn string
	if strings.HasPrefix(path, "/.clvq/") {
		fn = filepath.Join("/opt/src/site/static", strings.TrimPrefix(path, "/.clvq/"))
	} else {
		fn = filepath.Join(cfg.StaticDir(), path)
	}
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		http.Error(w, "404 - not found", http.StatusNotFound)
		log.Printf("404 %s - %v", path, err)
		return
	}
	ext := filepath.Ext(fn)
	mimeType := mime.TypeByExtension(ext)
	w.Header().Set("Content-Type", mimeType)
	http.ServeFile(w, r, fn)
	log.Printf("200 %s - %s", path, fn)
}

func (h *Handler) ServeTpl(w http.ResponseWriter, r *http.Request, path string) {
	tmplt, err := h.tpl.Get(path)
	if err != nil {
		if _, err := os.Stat(h.tpl.Filepath(path)); os.IsNotExist(err) {
			http.Error(w, "404 - not found", http.StatusNotFound)
			log.Printf("404 %s - %v", path, err)
			return
		}
	}
	data := h.tpl.GetData(path)
	if err := tmplt.Execute(w, data); err != nil {
		http.Error(w, "500 - failed to render template", http.StatusInternalServerError)
		log.Printf("500 %s - %v", path, err)
		return
	}
	log.Printf("200 %s - %s %s", path, h.tpl.BaseFile(), h.tpl.Filepath(path))
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	reqPath := path.Clean(r.URL.Path)
	if strings.HasSuffix(reqPath, "/") {
		reqPath = path.Join(reqPath, "index.html")
	}
	ext := path.Ext(reqPath)
	if ext == "" {
		reqPath = path.Join(reqPath, "index.html")
		ext = path.Ext(reqPath)
	}
	if ext == ".html" {
		h.ServeTpl(w, r, reqPath)
	} else {
		h.ServeFile(w, r, reqPath)
	}
}

func Main(port string) error {
	log.Printf("starting http server on port: %s", port)
	log.Printf("static dir: %s", cfg.StaticDir())
	log.Printf("tpl dir: %s", cfg.TplDir())

	http.HandleFunc("/clvq/", site.ServeStaticFS)
	http.HandleFunc("/_/config.json", admin.ConfigJSONHandler)
	AddHandler("/_/", admin.NewTpl())

	log.Printf("admin tpl dev: %v", env.AdminTplDev)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}
