// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package build

import (
	"log"

	"chesslovaquia.github.io/go/clvq/http"
)

func Main() error {
	log.SetFlags(0)
	http.Get("/")
	return nil
}
