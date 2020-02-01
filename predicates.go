package plusminus

import (
	"fmt"
	"strings"
)

// Predicate defines a predicate to return from a node.
func Predicate(name string) *predicate {
	return &predicate{
		name: name,
	}
}

type predicateList []*predicate

func (p predicateList) toString() string {
	var predicatesGraphQLpm string
	for i := 0; i < len(p); i++ {
		predicatesGraphQLpm += p[i].toString() + "\n"
	}

	return "{\n" + predicatesGraphQLpm + "}"
}

type predicate struct {
	name        string
	filter      expr
	first       interface{}
	offset      interface{}
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

func (p *predicate) Filter(filter expr) *predicate {
	p.filter = filter
	return p
}

func (p *predicate) First(val interface{}) *predicate {
	p.first = val
	return p
}

func (p *predicate) Offset(val interface{}) *predicate {
	p.offset = val
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
	s := p.name

	if p.first != nil || p.offset != nil {
		s += " ("

		var pagination []string

		if p.first != nil {
			switch p.first.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				pagination = append(pagination, fmt.Sprintf("first: %d", p.first))
			default:
				pagination = append(pagination, fmt.Sprintf("first: %s", p.first))
			}
		}

		if p.offset != nil {
			switch p.offset.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				pagination = append(pagination, fmt.Sprintf("offset: %d", p.offset))
			default:
				pagination = append(pagination, fmt.Sprintf("offset: %s", p.offset))
			}
		}

		s += strings.Join(pagination, ", ")

		s += ")"
	}

	if p.filter != nil {
		s += " @filter(" + p.filter.toString() + ")"
	}

	if p.facets {
		if len(p.facetNames) == 0 && p.facetFilter == nil {
			s += " @facets"
		} else {
			if p.facetFilter != nil {
				s += " @facets(" + p.facetFilter.toString() + ")"
			}
			if len(p.facetNames) > 0 {
				s += " @facets(" + strings.Join(p.facetNames, ", ") + ")"
			}
		}
	}
	if len(p.predicates) > 0 {
		s += " " + p.predicates.toString()
	}

	return s
}
