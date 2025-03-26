// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cfg

type Config struct {
	path string
}

var cfg *Config

func init() {
	cfg = newConfig()
}

func newConfig() *Config {
	return &Config{}
}

func Load(path string) error {
	cfg.path = path
	return nil
}
