[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)

# plusminus

[![Build Status](https://github.com/nicklanng/plusminus/workflows/CI/badge.svg)](https://github.com/nicklanng/plusminus/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/nicklanng/plusminus)](https://goreportcard.com/report/github.com/nicklanng/plusminus)
[![Documentation](https://godoc.org/github.com/nicklanng/plusminus?status.svg)](http://godoc.org/github.com/nicklanng/plusminus)
[![GitHub issues](https://img.shields.io/github/issues/nicklanng/plusminus.svg)](https://github.com/nicklanng/plusminus/issues)
[![license](https://img.shields.io/github/license/nicklanng/plusminus.svg?maxAge=2592000)](https://github.com/nicklanng/plusminus/LICENSE)
[![Release](https://img.shields.io/github/release/nicklanng/plusminus.svg?label=Release)](https://github.com/nicklanng/plusminus/releases)

A query builder for [dgraph](https://dgraph.io/).

```go
import "github.com/nicklanng/plusminus"
```

PlusMinus helps to compose query statements for dgraph, especially in cases that queries must be built bit by bit.

The library only creates a query string, which can be then passed to something like [github.com/dgraph-io/dgo](https://github.com/dgraph-io/dgo).

```go
import pm "github.com/nicklanng/plusminus"

q := pm.Query("queryfriends").Blocks(
  pm.Block("data", pm.Eq("name", "Alice")).Predicates(
    pm.Predicate("name"),
    pm.Predicate("friend").Facets("close").Predicates(
      pm.Predicate("name"),
      pm.Predicate("car").Facets(),
    ),
  ),
).ToString()
}
```
