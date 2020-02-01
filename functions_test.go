package plusminus

import (
	"testing"
)

func Test_funcUID_toString(t *testing.T) {
	cases := map[string]struct {
		value    interface{}
		expected string
	}{
		"string value": {
			value:    "$1",
			expected: "uid($1)",
		},
		"int value": {
			value:    uint64(17),
			expected: "uid(0x11)",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			res := funcUID{c.value}.toString()
			if res != c.expected {
				t.Errorf("\nexpected: %s\ngot: %s", c.expected, res)
			}
		})
	}
}

func Test_funcEq_toString(t *testing.T) {
	cases := map[string]struct {
		name     string
		value    interface{}
		expected string
	}{
		"variable value": {
			name:     "name",
			value:    "$name",
			expected: `eq(name, $name)`,
		},
		"string value": {
			name:     "name",
			value:    "string",
			expected: `eq(name, "string")`,
		},
		"int value": {
			name:     "name",
			value:    54,
			expected: `eq(name, 54)`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			res := funcEq{c.name, c.value}.toString()
			if res != c.expected {
				t.Errorf("\nexpected: %s\ngot: %s", c.expected, res)
			}
		})
	}
}
