// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package http

import (
	"log"
	"net/http"

	_ "chesslovaquia.github.io/go/clvq/cfg"
)

func Main(port string) error {
	log.Printf("starting http server on port: %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}
