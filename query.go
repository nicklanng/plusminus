package plusminus

// Query is the top level structure of a dgraph query.
func Query(name string, filterTerm expr) *query {
	return &query{
		name:       name,
		filterTerm: filterTerm,
	}
}

type query struct {
	name       string
	filterTerm expr
	normalize  bool
	predicates predicateList
}

// Normalize will toggle on normalization of predicate and facet results.
// Only predicates and facets with an alias will be returned.
func (q *query) Normalize() *query {
	q.normalize = true
	return q
}

// Predicates is the list of predicates to return from the current node.
// Predicates can be nested.
func (q *query) Predicates(preds ...*predicate) *query {
	q.predicates = append(q.predicates, preds...)
	return q
}

func (q query) ToString() string {
	return q.name + "(func: " + q.filterTerm.toString() + ") " + q.predicates.toString() + "\n"
}
