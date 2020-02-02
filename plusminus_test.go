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

func TestDataQueryWithFacets(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("data", pm.Eq("name", "Alice")).Predicates(
			pm.Predicate("name"),
			pm.Predicate("mobile").Facets(),
			pm.Predicate("car").Facets(),
		),
	)

	expected := `query {
data(func: eq(name, "Alice")) {
name
mobile @facets
car @facets
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestDataQueryWithFacetsSelectors(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("data", pm.Eq("name", "Alice")).Predicates(
			pm.Predicate("name"),
			pm.Predicate("mobile").Facets("since"),
			pm.Predicate("car").Facets("since"),
		),
	)

	expected := `query {
data(func: eq(name, "Alice")) {
name
mobile @facets(since)
car @facets(since)
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestDataQueryWithFacetsAliases(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("data", pm.Eq("name", "Alice")).Predicates(
			pm.Predicate("name"),
			pm.Predicate("mobile"),
			pm.Predicate("car").Facets("car_since: since"),
			pm.Predicate("friend").Facets("close_friend: close"),
		),
	)

	expected := `query {
data(func: eq(name, "Alice")) {
name
mobile
car @facets(car_since: since)
friend @facets(close_friend: close)
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestDataQueryWithFacetsInUID(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("data", pm.Eq("name", "Alice")).Predicates(
			pm.Predicate("name"),
			pm.Predicate("friend").Facets("close").Predicates(
				pm.Predicate("name"),
				pm.Predicate("car").Facets(),
			),
		),
	)

	expected := `query {
data(func: eq(name, "Alice")) {
name
friend @facets(close) {
name
car @facets
}
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestDataQueryWithFacetsFilterWithAnd(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("data", pm.Eq("name", "Alice")).Predicates(
			pm.Predicate("friend").FacetFilter(pm.And(pm.Eq("close", true), pm.Eq("relative", true))).Facets("relative").Predicates(
				pm.Predicate("name"),
			),
		),
	)

	expected := `query {
data(func: eq(name, "Alice")) {
friend @facets(eq(close, true) AND eq(relative, true)) @facets(relative) {
name
}
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestMeQueryWithPaginationAndVariables(t *testing.T) {
	q := pm.Query("test").Variables("$b: int", "$name: string").Blocks(
		pm.Block("me", pm.AllOfTerms("name@en", "$name")).Predicates(
			pm.Predicate("name@en"),
			pm.Predicate("director.film").First(2).Offset("$b").Predicates(
				pm.Predicate("name@en"),
				pm.Predicate("genre").First("$a").Predicates(
					pm.Predicate("name@en"),
				),
			),
		),
	)

	expected := `query test($b: int, $name: string) {
me(func: allofterms(name@en, $name)) {
name@en
director.film(first: 2, offset: $b) {
name@en
genre(first: $a) {
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

func TestDirectorQueryWithNormalization(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("director", pm.AllOfTerms("name@en", "steven spielberg")).Normalize().Predicates(
			pm.Predicate("director: name@en"),
			pm.Predicate("director.film").Predicates(
				pm.Predicate("film: name@en"),
				pm.Predicate("initial_release_date"),
				pm.Predicate("starring").First(2).Predicates(
					pm.Predicate("performance.actor").Predicates(
						pm.Predicate("actor: name@en"),
					),
					pm.Predicate("performance.character").Predicates(
						pm.Predicate("character: name@en"),
					),
				),
				pm.Predicate("country").Predicates(
					pm.Predicate("country: name@en"),
				),
			),
		),
	)

	expected := `query {
director(func: allofterms(name@en, "steven spielberg")) @normalize {
director: name@en
director.film {
film: name@en
initial_release_date
starring(first: 2) {
performance.actor {
actor: name@en
}
performance.character {
character: name@en
}
}
country {
country: name@en
}
}
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}

func TestDirectorQueryWithNestedNormalization(t *testing.T) {
	q := pm.Query("").Blocks(
		pm.Block("director", pm.AllOfTerms("name@en", "steven spielberg")).Predicates(
			pm.Predicate("director: name@en"),
			pm.Predicate("director.film").Predicates(
				pm.Predicate("film: name@en"),
				pm.Predicate("initial_release_date"),
				pm.Predicate("starring").First(2).Normalize().Predicates(
					pm.Predicate("performance.actor").Predicates(
						pm.Predicate("actor: name@en"),
					),
					pm.Predicate("performance.character").Predicates(
						pm.Predicate("character: name@en"),
					),
				),
				pm.Predicate("country").Predicates(
					pm.Predicate("country: name@en"),
				),
			),
		),
	)

	expected := `query {
director(func: allofterms(name@en, "steven spielberg")) {
director: name@en
director.film {
film: name@en
initial_release_date
starring(first: 2) @normalize {
performance.actor {
actor: name@en
}
performance.character {
character: name@en
}
}
country {
country: name@en
}
}
}
}`

	result := q.ToString()
	if result != expected {
		t.Errorf("\nexpected: %s\ngot: %s", expected, result)
	}
}
