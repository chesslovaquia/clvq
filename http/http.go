// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package http

import (
	"log"
	"net/http"
	"path"
	"strings"

	"chesslovaquia.github.io/go/clvq/admin"
	"chesslovaquia.github.io/go/clvq/tpl"
)

var handlers map[string]*Handler

func init() {
	handlers = make(map[string]*Handler)
}

func AddHandler(path string, template tpl.Tpl) {
	handlers[path] = newHandler(template)
	http.HandleFunc(path, handlers[path].Handle)
}

func Get(path string) (tpl.Tpl, error) {
	return nil, nil
}

func ServeTpl(w http.ResponseWriter, r *http.Request, path string) {
	Get(path)
}

func ServeFile(w http.ResponseWriter, r *http.Request, path string) {
}

type Handler struct {
	template tpl.Tpl
}

func newHandler(template tpl.Tpl) *Handler {
	return &Handler{
		template: template,
	}
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
		ServeTpl(w, r, reqPath)
	} else {
		ServeFile(w, r, reqPath)
	}
}

func Main(port string) error {
	log.Printf("starting http server on port: %s", port)

	http.Handle("/_/static/", http.StripPrefix("/_/", http.FileServer(http.FS(admin.StaticFS))))
	AddHandler("/_/", admin.NewTpl())

	AddHandler("/", tpl.New())

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}
