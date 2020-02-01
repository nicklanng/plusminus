package plusminus

import "fmt"

// UID is a function that selects a node with the provided uid.
func UID(val uint64) funcUID {
	return funcUID{val: val}
}

type funcUID struct {
	val uint64
}

func (p funcUID) toString() string {
	return fmt.Sprintf("uid(%d)", p.val) + ")"
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
	switch p.val.(type) {
	case string:
		return fmt.Sprintf("eq(%s, %q)", p.name, p.val)
	default:
		return fmt.Sprintf("eq(%s, %v)", p.name, p.val)
	}
}
