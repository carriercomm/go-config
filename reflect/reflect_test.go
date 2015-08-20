package reflect_test

import (
	"testing"

	"github.com/ttaylorr/go-config/reflect"
)

func TestShallowMapLookups(t *testing.T) {
	root := map[string]interface{}{
		"foo": "bar",
	}
	assertFetch("foo", "bar", root, t)
}

func TestShallowArrayLookups(t *testing.T) {
	root := []interface{}{
		"foo", "bar",
	}
	assertFetch("1", "bar", root, t)
}

func TestNestedMapLookups(t *testing.T) {
	root := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"baz": "donk",
			},
		},
	}
	assertFetch("foo.bar.baz", "donk", root, t)
}

func TestNestedArrayLookups(t *testing.T) {
	root := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": []interface{}{
				"woot", "boop", "bop", "donk", "flub",
			},
		},
	}
	assertFetch("foo.bar.3", "donk", root, t)
}

func assertFetch(lookup string, expected interface{}, root interface{}, t *testing.T) {
	if val, err := reflect.Fetch(lookup, root); err != nil {
		t.Errorf("expected no error (got %v)", err)
	} else if val != expected {
		t.Errorf("expected value to equal \"%v\" (got \"%v\")", expected, val)
	}
}
