package plusminus

import (
	"testing"
)

func TestFilterTermEq_toString(t *testing.T) {
	cases := map[string]struct {
		name     string
		value    interface{}
		expected string
	}{
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

	for _, c := range cases {
		res := Eq(c.name, c.value).toString()
		if res != c.expected {
			t.Errorf("\nexpected: %s\ngot: %s", c.expected, res)
		}
	}
}
