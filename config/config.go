package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ttaylorr/go-config/environment"
	"github.com/ttaylorr/go-config/reflect"
)

type Configuration struct {
	root   interface{}
	parent *Configuration
}

func New(configDir string) (*Configuration, error) {
	return inew(configDir, environment.Current())
}

func inew(configDir string, env environment.Environment) (*Configuration, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s.json", configDir, string(env)))
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	var root interface{}

	if err := decoder.Decode(&root); err != nil {
		return nil, err
	}

	var parent *Configuration
	var err error
	if env.Parent() != env {
		parent, err = inew(configDir, env.Parent())
		if err != nil {
			return nil, err
		}
	}

	return &Configuration{root, parent}
}

func (c *Configuration) get(lookup string) (interface{}, error) {
	i, err := reflect.Fetch(lookup, c.root)

	if err != nil {
		if c.parent != nil {
			return c.parent.get(lookup)
		}

		return nil, err
	}

	return i, nil
}
