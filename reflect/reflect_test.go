package reflect_test

import (
	"testing"

	"github.com/ttaylorr/go-config/reflect"
)

func TestShallowMapLookups(t *testing.T) {
	assertFetch("foo", "bar", map[string]interface{}{
		"foo": "bar",
	}, t)
}

func TestShallowArrayLookups(t *testing.T) {
	assertFetch("1", "bar", []interface{}{
		"foo", "bar",
	}, t)
}

func TestNestedMapLookups(t *testing.T) {
	assertFetch("foo.bar.baz", "donk", map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"baz": "donk",
			},
		},
	}, t)
}

func TestNestedArrayLookups(t *testing.T) {
	assertFetch("foo.bar.3", "donk", map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": []interface{}{
				"woot", "boop", "bop", "donk", "flub",
			},
		},
	}, t)
}

func assertFetch(lookup string, expected interface{}, root interface{}, t *testing.T) {
	if val, err := reflect.Fetch(lookup, root); err != nil {
		t.Errorf("expected no error (got %v)", err)
	} else if val != expected {
		t.Errorf("expected value to equal \"%v\" (got \"%v\")", expected, val)
	}
}
