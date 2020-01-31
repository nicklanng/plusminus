package plusminus

import "fmt"

// UID is a filter expression that selects a node with the provided uid.
func UID(val uint64) filterTermUID {
	return filterTermUID{val: val}
}

type filterTermUID struct {
	val uint64
}

func (p filterTermUID) toString() string {
	return fmt.Sprintf("uid(%d)", p.val) + ")"
}

// Eq is a filter expression that selects nodes with the matching value.
func Eq(name string, val interface{}) filterTermEq {
	return filterTermEq{name: name, val: val}
}

type filterTermEq struct {
	name string
	val  interface{}
}

func (p filterTermEq) toString() string {
	return fmt.Sprintf("eq(%s, %v)", p.name, p.val)
}
