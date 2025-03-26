// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package clvq

import (
	"flag"
	"log"

	"chesslovaquia.github.io/go/clvq/cfg"
)

var die func(string, ...any)

func init() {
	die = log.Fatalf
}

func Main() {
	configFilename := flag.String("config", "clvq.json", "config filename")
	flag.Parse()

	if err := cfg.Load(*configFilename); err != nil {
		die("[ERROR] load config: %v", err)
	}
}
