package plusminus

import "strings"

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

// ToString creates a string representation of the whole query, which can be used to query dgraph.
func (q *query) ToString() string {
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
