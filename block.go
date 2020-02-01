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
	filter     expr
	predicates predicateList
}

// Normalize will toggle on normalization of predicate and facet results.
// Only predicates and facets with an alias will be returned.
func (b *block) Normalize() *block {
	b.normalize = true
	return b
}

func (b *block) Filter(filter expr) *block {
	b.filter = filter
	return b
}

// Predicates allows you to add a number of predicates to the current returned node.
// Predicates can be nested.
func (b *block) Predicates(preds ...*predicate) *block {
	b.predicates = append(b.predicates, preds...)
	return b
}

func (b block) toString() string {
	s := b.name + "(func: " + b.function.toString() + ") "

	if b.filter != nil {
		s += "@filter(" + b.filter.toString() + ") "
	}

	s += b.predicates.toString() + "\n"

	return s
}
