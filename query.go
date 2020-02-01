package plusminus

// Query is a top-level query to dgraph. It can include mulitple query blocks.
func Query() *query {
	return &query{}
}

type query struct {
	blocks []*block
}

func (q *query) Blocks(b ...*block) *query {
	q.blocks = append(q.blocks, b...)
	return q
}

func (q *query) ToString() string {
	s := "query "

	s += "{\n"
	for i := range q.blocks {
		s += q.blocks[i].toString()
	}
	s += "}"

	return s
}
