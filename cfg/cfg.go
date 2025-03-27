// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

import (
	"encoding/json"
	"os"
)

// tpl

var Tpl *ConfigTpl

type ConfigTpl struct{
	Dir string
}

func newConfigTpl() *ConfigTpl {
	return &ConfigTpl{
		Dir: "tpl",
	}
}

// cfg

type Config struct {
	path string
	Tpl *ConfigTpl
}

var cfg *Config

func init() {
	Tpl = newConfigTpl()
	cfg = newConfig()
}

func newConfig() *Config {
	return &Config{
		Tpl: Tpl,
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
