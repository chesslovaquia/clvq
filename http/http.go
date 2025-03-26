// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package http

import (
	"log"
	"net/http"
	"path"
	"strings"

	"chesslovaquia.github.io/go/clvq/tpl"
)

func Get(path string) (tpl.Tpl, error) {
	return nil, nil
}

func ServeTpl(w http.ResponseWriter, r *http.Request, path string) {
	Get(path)
}

func ServeFile(w http.ResponseWriter, r *http.Request, path string) {
}

func Handler(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/", Handler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}
