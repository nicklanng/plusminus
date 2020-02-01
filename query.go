package plusminus

// Query is a top-level query to dgraph. It can include mulitple query blocks.
func Query(blocks ...*block) query {
	return query{
		blocks: blocks,
	}
}

type query struct {
	blocks []*block
}

func (q query) ToString() string {
	s := "query "

	s += "{\n"
	for i := range q.blocks {
		s += q.blocks[i].toString()
	}
	s += "}"

	return s
}
