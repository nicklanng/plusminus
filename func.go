package plusminus

// Func is the top level statement of a query. It accepts a filter term.
func Func(filterTerm expr) *funcStmt {
	return &funcStmt{filterTerm: filterTerm}
}

type funcStmt struct {
	filterTerm expr
	normalize  bool
	predicates predicateList
}

// Normalize will toggle on normalization of predicate and facet results.
// Only predicates and facets with an alias will be returned.
func (p *funcStmt) Normalize() *funcStmt {
	p.normalize = true
	return p
}

// Predicates is the list of predicates to return from the current node.
// Predicates can be nested.
func (p *funcStmt) Predicates(preds ...*predicate) *funcStmt {
	p.predicates = append(p.predicates, preds...)
	return p
}

func (p funcStmt) toString() string {
	return "q(func: " + p.filterTerm.toString() + p.predicates.toString()
}
