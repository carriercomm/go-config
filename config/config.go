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

type Param struct {
	// Directory refers to the directory in which to search for
	// configuration files. Can be either relative to the current path of
	// execution, or absolute.
	Directory string

	// Environment refers to the environment for which to load
	// configuration files. If this field is not specified, the `None`
	// Environment const, will be passed, and promoted up to `Default`.
	Environment environment.Environment
}

func New(p *Param) (*Configuration, error) {
	e := p.Environment
	if e == environment.None {
		e = environment.Default
	}

	return inew(p.Directory, e)
}

// TODO(ttaylorr): clean
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
	var perr error
	if env.Parent() != env {
		parent, perr = inew(configDir, env.Parent())
		if perr != nil {
			return nil, perr
		}
	}

	return &Configuration{root, parent}, nil
}

func (c *Configuration) String(lookup string) (string, error) {
	if v, err := c.get(lookup); err != nil {
		return "", err
	} else {
		if s, err := reflect.Coerse(v, reflect.String); err != nil {
			return "", err
		} else {
			return s.(string), nil
		}
	}
}

func (c *Configuration) Int(lookup string) (int, error) {
	if v, err := c.get(lookup); err != nil {
		return -1, err
	} else {
		if i, err := reflect.Coerse(v, reflect.Int); err != nil {
			return -1, err
		} else {
			return i.(int), nil
		}
	}
}

func (c *Configuration) Bool(lookup string) (bool, error) {
	if v, err := c.get(lookup); err != nil {
		return false, err
	} else {
		if b, err := reflect.Coerse(v, reflect.Bool); err != nil {
			return false, err
		} else {
			return b.(bool), nil
		}
	}
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
