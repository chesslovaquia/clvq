// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package clvq

import (
	"flag"
	"log"

	"chesslovaquia.github.io/go/clvq/build"
	"chesslovaquia.github.io/go/clvq/cfg"
	"chesslovaquia.github.io/go/clvq/http"
	"chesslovaquia.github.io/go/clvq/tpl"
)

var die func(string, ...any)

func init() {
	die = log.Fatalf
}

func Main() {
	configFilename := flag.String("config", "clvq.json", "config filename")
	httpPort := flag.String("http", "8044", "http port")

	runBuild := flag.Bool("build", false, "run site build")

	flag.Parse()

	if err := cfg.Load(*configFilename); err != nil {
		die("[ERROR] load config: %v", err)
	}

	if *runBuild {
		if err := build.Main(); err != nil {
			die("[ERROR] build site: %v", err)
		}
		return
	}

	if err := http.Main(*httpPort); err != nil {
		die("[ERROR] http server: %v", err)
	}
}

func AddHandler(path string, template tpl.Tpl) {
	http.AddHandler(path, template)
}
