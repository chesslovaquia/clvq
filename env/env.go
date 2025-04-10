// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package env

import (
	"os"
	"strings"
)

func getBool(name string, defval bool) bool {
	val := os.Getenv(name)
	if val != "" {
		switch strings.ToLower(strings.TrimSpace(val)) {
		case "1", "t", "true", "yes", "y", "enable":
			return true
		default:
			return false
		}
	}
	return defval
}

func get(name, defval string) string {
	val := os.Getenv(name)
	if val != "" {
		return strings.TrimSpace(val)
	}
	return defval
}

var (
	AdminTplDev bool
)

func init() {
	AdminTplDev = getBool("CLVQ_ADMIN_TPL_DEV", false)
}
