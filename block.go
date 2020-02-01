package plusminus

// Block is a query block that can make up part of a query.
func Block(name string, function expr) *block {
	return &block{
		name:     name,
		function: function,
	}
}

type block struct {
	name       string
	function   expr
	normalize  bool
	predicates predicateList
}

// Normalize will toggle on normalization of predicate and facet results.
// Only predicates and facets with an alias will be returned.
func (q *block) Normalize() *block {
	q.normalize = true
	return q
}

// Predicates allows you to add a number of predicates to the current returned node.
// Predicates can be nested.
func (q *block) Predicates(preds ...*predicate) *block {
	q.predicates = append(q.predicates, preds...)
	return q
}

func (q block) toString() string {
	return q.name + "(func: " + q.function.toString() + ") " + q.predicates.toString() + "\n"
}
