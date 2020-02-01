package plusminus

type expr interface {
	toString() string
}

// Query is the top level structure of a dgraph query.
func Query(funcs ...*funcStmt) query {
	return query{
		funcs: funcs,
	}
}

type query struct {
	funcs []*funcStmt
}

func (q query) ToString() string {
	s := "{\n"
	for i := range q.funcs {
		s += q.funcs[i].toString()
	}
	s += "}\n"
	return s
}
