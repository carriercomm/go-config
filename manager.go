package config

import (
	"fmt"
	"os"
	"strings"
)

// XXX(ttaylorr): global state
var configs map[string]*Configuration = make(map[string]*Configuration, 0)

func Get(dir string) (*Configuration, error) {
	path := envPath(dir)

	if configs[path] == nil {
		if cfg, err := newConfig(path); err != nil {
			return nil, err
		} else {
			// TODO(ttaylorr): attach parent here
			configs[path] = cfg
		}
	}

	return configs[path]
}

func envPath(dir string) string {
	return fmt.Sprintf("%s/%s.json", dir, Environment())
}

func Environment() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "development"
	}

	return strings.ToLower(env)
}
