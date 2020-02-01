package plusminus

import "strings"

// Predicate defines a predicate to return from a node.
func Predicate(name string) *predicate {
	return &predicate{
		name: name,
	}
}

type predicateList []*predicate

func (p predicateList) toString() string {
	if len(p) == 0 {
		return ""
	}

	var predicatesGraphQLpm string
	for i := 0; i < len(p); i++ {
		predicatesGraphQLpm += p[i].toString() + "\n"
	}

	return "{\n" + predicatesGraphQLpm + "}"
}

type predicate struct {
	name        string
	facets      bool
	facetNames  []string
	facetFilter expr
	predicates  predicateList
}

// Predicates allows you to add a number of predicates to the current returned node.
// Predicates can be nested.
func (p *predicate) Predicates(preds ...*predicate) *predicate {
	p.predicates = append(p.predicates, preds...)
	return p
}

func (p *predicate) Facets(names ...string) *predicate {
	p.facets = true
	p.facetNames = names
	return p
}
func (p *predicate) FacetFilter(filter expr) *predicate {
	p.facets = true
	p.facetFilter = filter
	return p
}

func (p *predicate) toString() string {
	s := p.name + " "
	if p.facets {
		if len(p.facetNames) == 0 && p.facetFilter == nil {
			s += "@facets "
		} else {
			if p.facetFilter != nil {
				s += "@facets(" + p.facetFilter.toString() + ") "
			}
			if len(p.facetNames) > 0 {
				s += "@facets(" + strings.Join(p.facetNames, ", ") + ") "
			}
		}
	}
	s += p.predicates.toString()
	return s
}
