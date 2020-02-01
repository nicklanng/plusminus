package plusminus

import (
	"testing"
)

func Test_funcUID_toString(t *testing.T) {
	cases := map[string]struct {
		values   []interface{}
		expected string
	}{
		"variable value": {
			values:   []interface{}{"$1"},
			expected: "uid($1)",
		},
		"int value": {
			values:   []interface{}{uint64(17)},
			expected: "uid(0x11)",
		},
		"mulitple int values": {
			values:   []interface{}{uint64(17), uint64(18)},
			expected: "uid(0x11, 0x12)",
		},
		"mulitple variable values": {
			values:   []interface{}{"$1", "$2"},
			expected: "uid($1, $2)",
		},
		"mulitple variable and int values": {
			values:   []interface{}{"$1", uint64(17), "$2"},
			expected: "uid($1, 0x11, $2)",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := funcUID{c.values}.toString()
			if result != c.expected {
				t.Errorf("\nexpected: %s\ngot: %s", c.expected, result)
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
			result := funcEq{c.name, c.value}.toString()
			if result != c.expected {
				t.Errorf("\nexpected: %s\ngot: %s", c.expected, result)
			}
		})
	}
}

func Test_funcLe_toString(t *testing.T) {
	cases := map[string]struct {
		name     string
		value    interface{}
		expected string
	}{
		"variable value": {
			name:     "name",
			value:    "$name",
			expected: `le(name, $name)`,
		},
		"string value": {
			name:     "name",
			value:    "string",
			expected: `le(name, "string")`,
		},
		"int value": {
			name:     "name",
			value:    54,
			expected: `le(name, 54)`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := funcLe{c.name, c.value}.toString()
			if result != c.expected {
				t.Errorf("\nexpected: %s\ngot: %s", c.expected, result)
			}
		})
	}
}

func Test_funcHas_toString(t *testing.T) {
	cases := map[string]struct {
		name     string
		expected string
	}{
		"string value": {
			name:     "name",
			expected: `has(name)`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := funcHas{c.name}.toString()
			if result != c.expected {
				t.Errorf("\nexpected: %s\ngot: %s", c.expected, result)
			}
		})
	}
}

func Test_funcAllOfTerms_toString(t *testing.T) {
	cases := map[string]struct {
		name     string
		value    string
		expected string
	}{
		"string value": {
			name:     "name",
			value:    "string second",
			expected: `allofterms(name, "string second")`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := funcAllOfTerms{c.name, c.value}.toString()
			if result != c.expected {
				t.Errorf("\nexpected: %s\ngot: %s", c.expected, result)
			}
		})
	}
}
