package plusminus

// Query is a top-level query to dgraph.
func Query() *query {
	return &query{}
}

type query struct {
	blocks []*block
}

// Blocks allows you to add a number of blocks to the query.
func (q *query) Blocks(b ...*block) *query {
	q.blocks = append(q.blocks, b...)
	return q
}

// ToString creates a string representation of the whole query, which can be used to query dgraph.
func (q *query) ToString() string {
	s := "query "

	s += "{\n"
	for i := range q.blocks {
		s += q.blocks[i].toString()
	}
	s += "}"

	return s
}
