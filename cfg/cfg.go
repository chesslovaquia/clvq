// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"encoding/json"
	"os"
)

// tpl

var Tpl *TplData

type TplData struct {
	Root string
	Site string
}

func newTplData() *TplData {
	return &TplData{
		Root: "",
		Site: "clvq",
	}
}

// cfg

type Config struct {
	path      string
	Tpl       *TplData
	StaticDir string
	TplDir    string
	TplBase   string
}

var cfg *Config

func init() {
	Tpl = newTplData()
	cfg = newConfig()
}

func newConfig() *Config {
	return &Config{
		Tpl:       Tpl,
		StaticDir: "static",
		TplDir:    "tpl",
		TplBase:   "base.html",
	}
}

func Save(path string) error {
	fh, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fh.Close()
	encoder := json.NewEncoder(fh)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cfg); err != nil {
		return err
	}
	return nil
}

var autoSave bool = true

func Load(path string) error {
	if blob, err := os.ReadFile(path); err != nil {
		if os.IsNotExist(err) {
			if autoSave {
				autoSave = false
				if err := Save(path); err != nil {
					return err
				} else {
					return Load(path)
				}
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		if err := json.Unmarshal(blob, cfg); err != nil {
			return err
		}
	}
	cfg.path = path
	return nil
}

func StaticDir() string {
	return cfg.StaticDir
}

func TplDir() string {
	return cfg.TplDir
}

func TplBase() string {
	return cfg.TplBase
}
