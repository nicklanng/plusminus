package plusminus_test

import (
	"testing"

	pm "github.com/nicklanng/plusminus"
)

func TestBladerunnerQuery(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("bladerunner", pm.Eq("name@en", "Blade Runner")).Predicates(
			pm.Predicate("uid"),
			pm.Predicate("name@en"),
			pm.Predicate("initial_release_date"),
			pm.Predicate("netflix_id"),
		),
	)

	expected := `query {
bladerunner(func: eq(name@en, "Blade Runner")) {
uid
name@en
initial_release_date
netflix_id
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestMoviesQueryWithMultipleUIDs(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("movies", pm.UID(0xb5849, 0x394c)).Predicates(
			pm.Predicate("uid"),
			pm.Predicate("name@en"),
			pm.Predicate("initial_release_date"),
			pm.Predicate("netflix_id"),
		),
	)

	expected := `query {
movies(func: uid(0xb5849, 0x394c)) {
uid
name@en
initial_release_date
netflix_id
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestBRCharactersQueryWithNestedPredicates(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("brCharacters", pm.Eq("name@en", "Blade Runner")).Predicates(
			pm.Predicate("name@en"),
			pm.Predicate("initial_release_date"),
			pm.Predicate("starring").Predicates(
				pm.Predicate("performance.actor").Predicates(
					pm.Predicate("name@en"),
				),
				pm.Predicate("performance.character").Predicates(
					pm.Predicate("name@en"),
				),
			),
		),
	)

	expected := `query {
brCharacters(func: eq(name@en, "Blade Runner")) {
name@en
initial_release_date
starring {
performance.actor {
name@en
}
performance.character {
name@en
}
}
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestScottQueryWithFilters(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("scott", pm.Eq("name@en", "Ridley Scott")).Predicates(
			pm.Predicate("name@en"),
			pm.Predicate("initial_release_date"),
			pm.Predicate("director.film").Filter(pm.Le("initial_release_date", "2000")).Predicates(
				pm.Predicate("name@en"),
				pm.Predicate("initial_release_date"),
			),
		),
	)

	expected := `query {
scott(func: eq(name@en, "Ridley Scott")) {
name@en
initial_release_date
director.film @filter(le(initial_release_date, "2000")) {
name@en
initial_release_date
}
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestMeQueryWithHasAndAllOfTermsFilter(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("me", pm.Eq("name@en", "Steven Spielberg")).Filter(pm.Has("director.film")).Predicates(
			pm.Predicate("name@en"),
			pm.Predicate("director.film").Filter(pm.AllOfTerms("name@en", "jones indiana")).Predicates(
				pm.Predicate("name@en"),
			),
		),
	)

	expected := `query {
me(func: eq(name@en, "Steven Spielberg")) @filter(has(director.film)) {
name@en
director.film @filter(allofterms(name@en, "jones indiana")) {
name@en
}
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}
