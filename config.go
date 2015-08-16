package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Configuration struct {
	root interface{}
	path string
}

func newConfig(dir string) (*Configuration, error) {
	path := envPath(dir)
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var root interface{}
	if err := json.Unmarshall(contents, &root); err != nil {
		return err
	}

	return &Configuration{root, path}, nil
}

func (c *Configuration) String(q string) (string, error) {
	e, err := c.get(q)
	if err != nil {
		return nil, err
	}

	switch t := e.(type) {
	case string:
		return t, nil
	default:
		return nil, errors.New("expected type string for query %s", q)
	}
}

func (c *Configuration) getFallback(q string) (interface{}, error) {
	var elem interface{}
	var err error

	elem, err = c.get(q)
	if err != nil && c.Parent != nil {
		elem, err = c.Parent.get(q)
	}

	if err != nil {
		return nil, err
	}
	return elem, nil
}

func (c *Configuration) get(q string) (interface{}, error) {
	namespaces := strings.Split(q, ".")

	elem := c.root
	for _, ns := range namespaces {
		switch t := elem.(type) {
		case []interface{}:
			if i, err := strconv.ParseInt(ns, 10, 0); err == nil {
				elem = elem[i]
			} else {
				return nil, fmt.Errorf("expected number (got %v)", ns)
			}
		case map[string]interface{}:
			if val, ok := elem[ns]; ok {
				elem = val
			} else {
				return nil, fmt.Errorf("no object with name %s", ns)
			}
		default:
			return nil, fmt.Errorf("invalid type found at namespace point (expected map[string]interface{}, or []interface{})")
		}
	}

	return elem, nil
}
