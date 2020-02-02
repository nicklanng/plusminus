package plusminus

import (
	"bytes"
	"strings"
)

// Query is a top-level query to dgraph.
func Query(name string) *query {
	return &query{
		name: name,
	}
}

type query struct {
	name      string
	variables []string
	blocks    []*block
}

// Variables allows you to add a number of variables to a query.
func (q *query) Variables(v ...string) *query {
	q.variables = append(q.variables, v...)
	return q
}

// Blocks allows you to add a number of blocks to the query.
func (q *query) Blocks(b ...*block) *query {
	q.blocks = append(q.blocks, b...)
	return q
}

// String creates a string representation of the whole query, which can be used to query dgraph.
func (q *query) String() string {
	s := "query " + q.name

	if len(q.variables) > 0 {
		s += "(" + strings.Join(q.variables, ", ") + ") "
	}

	s += "{\n"
	for i := range q.blocks {
		s += q.blocks[i].toString()
	}
	s += "}"

	return s
}

// StringIndented creates a query string like String(), but with appropriate indentation.
// This is not required for dgraph to parse the query and it is more computationally expensive, so use only when human readability of the query is desired.
func (q *query) StringIndented() string {
	var (
		buf    = new(bytes.Buffer)
		depth  = 0
		indent = []byte{' ', ' '}
	)

	str := q.String()

	for i := range str {
		if str[i] == '{' {
			depth++
		}

		if i+1 < len(str) && str[i+1] == '}' {
			depth--
		}

		buf.WriteByte(str[i])

		if str[i] == '\n' {
			buf.Write(bytes.Repeat(indent, depth))
		}

	}

	return buf.String()
}
