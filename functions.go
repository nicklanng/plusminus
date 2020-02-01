package plusminus

import "fmt"

// UID is a function that selects a node with the provided uid.
func UID(val interface{}) funcUID {
	return funcUID{val: val}
}

type funcUID struct {
	val interface{}
}

func (p funcUID) toString() string {
	switch p.val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("uid(0x%x)", p.val)
	default:
		return fmt.Sprintf("uid(%s)", p.val)
	}
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
