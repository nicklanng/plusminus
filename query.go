package plusminus

func Query(blocks ...*block) query {
	return query{
		blocks: blocks,
	}
}

type query struct {
	blocks []*block
}

func (q query) ToString() string {
	s := "{\n"

	for i := range q.blocks {
		s += q.blocks[i].toString()
	}

	s += "}"

	return s
}
