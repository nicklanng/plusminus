package plusminus

import (
	"fmt"
	"strings"
)

// UID is a function that selects a node with the provided uid.
func UID(values ...interface{}) funcUID {
	return funcUID{values: values}
}

type funcUID struct {
	values []interface{}
}

func (p funcUID) toString() string {
	valuesAsStrings := make([]string, len(p.values))

	for i := range p.values {
		switch p.values[i].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			valuesAsStrings[i] = fmt.Sprintf("0x%x", p.values[i])
		default:
			valuesAsStrings[i] = fmt.Sprintf("%s", p.values[i])
		}
	}

	return "uid(" + strings.Join(valuesAsStrings, ", ") + ")"
}

// Eq is a function that selects nodes with the matching value.
func Eq(name string, val interface{}) funcEq {
	return funcEq{name: name, val: val}
}

type funcEq struct {
	name string
	val  interface{}
}

func (p funcEq) toString() string {
	switch v := p.val.(type) {
	case string:
		if v[0] == '$' {
			return fmt.Sprintf("eq(%s, %s)", p.name, v)
		}
		return fmt.Sprintf("eq(%s, %q)", p.name, v)
	default:
		return fmt.Sprintf("eq(%s, %v)", p.name, v)
	}
}

// Le is a function that selects nodes with values of the named field that are less than the supplied value.
func Le(name string, val interface{}) funcLe {
	return funcLe{name: name, val: val}
}

type funcLe struct {
	name string
	val  interface{}
}

func (p funcLe) toString() string {
	switch v := p.val.(type) {
	case string:
		if v[0] == '$' {
			return fmt.Sprintf("le(%s, %s)", p.name, v)
		}
		return fmt.Sprintf("le(%s, %q)", p.name, v)
	default:
		return fmt.Sprintf("le(%s, %v)", p.name, v)
	}
}

// Has is a function that checks for existence of the named predicate.
func Has(name string) funcHas {
	return funcHas{name: name}
}

type funcHas struct {
	name string
}

func (p funcHas) toString() string {
	return fmt.Sprintf("has(%s)", p.name)
}

// AllOfTerms is a function that selects nodes with all matching terms of the named field.
func AllOfTerms(name string, terms string) funcAllOfTerms {
	return funcAllOfTerms{name: name, terms: terms}
}

type funcAllOfTerms struct {
	name  string
	terms string
}

func (p funcAllOfTerms) toString() string {
	return fmt.Sprintf("allofterms(%s, %q)", p.name, p.terms)
}

// And is a composite function that requires both the left and right side to be true.
func And(left, right expr) funcAnd {
	return funcAnd{left: left, right: right}
}

type funcAnd struct {
	left  expr
	right expr
}

func (p funcAnd) toString() string {
	return p.left.toString() + " AND " + p.right.toString()
}
