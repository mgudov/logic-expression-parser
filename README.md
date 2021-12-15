logic-expression-parser
====

[![Build Status](https://github.com/mgudov/logic-expression-parser/actions/workflows/test.yml/badge.svg)](https://github.com/mgudov/logic-expression-parser/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mgudov/logic-expression-parser)](https://goreportcard.com/report/github.com/mgudov/logic-expression-parser)
[![Codecov](https://codecov.io/gh/mgudov/logic-expression-parser/branch/master/graph/badge.svg?token=JMQMBEP2Z6)](https://codecov.io/gh/mgudov/logic-expression-parser)

This library provide generic boolean expression parser to go structures.

## Installation

    $ go get -u github.com/mgudov/logic-expression-parser

(optional) Run unit tests

    $ make test

(optional) Run benchmarks

    $ make bench

## Examples

```go
package main

import (
	"github.com/davecgh/go-spew/spew"
	lep "github.com/mgudov/logic-expression-parser"
)

func main() {
	expression := `a=false && b>=c && (d<1000 || e in [1,2,3])`

	result, err := lep.ParseExpression(expression)
	if err != nil {
		panic(err)
	}

	dump := spew.NewDefaultConfig()
	dump.DisablePointerAddresses = true
	dump.DisableMethods = true
	dump.Dump(result)
}
```

This library would parse the expression and return the following struct:

```
(*lep.AndX)({
  Conjuncts: ([]lep.Expression) (len=3 cap=4) {
    (*lep.EqualsX)({
      Param: (*lep.ParamX)({
        Name: (string) (len=1) "a"
      }),
      Value: (*lep.BooleanX)({
        Val: (bool) false
      })
    }),
    (*lep.GreaterThanEqualX)({
      Param: (*lep.ParamX)({
        Name: (string) (len=1) "b"
      }),
      Value: (*lep.ParamX)({
        Name: (string) (len=1) "c"
      })
    }),
    (*lep.OrX)({
      Disjunctions: ([]lep.Expression) (len=2 cap=2) {
        (*lep.LessThanX)({
          Param: (*lep.ParamX)({
            Name: (string) (len=1) "d"
          }),
          Value: (*lep.IntegerX)({
            Val: (int64) 1000
          })
        }),
        (*lep.InSliceX)({
          Param: (*lep.ParamX)({
            Name: (string) (len=1) "e"
          }),
          Slice: (*lep.SliceX)({
            Values: ([]lep.Value) (len=3 cap=4) {
              (*lep.IntegerX)({
                Val: (int64) 1
              }),
              (*lep.IntegerX)({
                Val: (int64) 2
              }),
              (*lep.IntegerX)({
                Val: (int64) 3
              })
            }
          })
        })
      }
    })
  }
})
```

Use can also create expression string from code:

```go
package main

import (
	"fmt"
	lep "github.com/mgudov/logic-expression-parser"
)

func main() {
	expression := lep.And(
		lep.Equals(lep.Param("a"), lep.Boolean(false)),
		lep.GreaterThanEqual(lep.Param("b"), lep.Param("c")),
		lep.Or(
			lep.LessThan(lep.Param("d"), lep.Integer(1000)),
			lep.InSlice(
				lep.Param("e"),
				lep.Slice(lep.Integer(1), lep.Integer(2), lep.Integer(3)),
			),
		),
	)

	fmt.Println(expression.String())
}
```

```
a=false && b>=c && (d<1000 || e in [1,2,3])
```

## Real life examples
<details>
  <summary>Create SQL query from expression string</summary>

```go
package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	sb "github.com/huandu/go-sqlbuilder"
	lep "github.com/mgudov/logic-expression-parser"
)

func traverse(sql *sb.SelectBuilder, expr lep.Expression) (string, error) {
	switch e := expr.(type) {
	default:
		return "", fmt.Errorf("not implemented: %T", e)
	case *lep.OrX:
		var args []string
		for _, disjunction := range e.Disjunctions {
			arg, err := traverse(sql, disjunction)
			if err != nil {
				return "", err
			}
			args = append(args, arg)
		}
		return sql.Or(args...), nil
	case *lep.AndX:
		var args []string
		for _, conjunct := range e.Conjuncts {
			arg, err := traverse(sql, conjunct)
			if err != nil {
				return "", err
			}
			args = append(args, arg)
		}
		return sql.And(args...), nil
	case *lep.EqualsX:
		value := e.Value.Value()
		if value == nil {
			return sql.IsNotNull(e.Param.String()), nil
		}
		return sql.Equal(e.Param.String(), value), nil
	case *lep.NotEqualsX:
		value := e.Value.Value()
		if value == nil {
			return sql.IsNotNull(e.Param.String()), nil
		}
		return sql.NotEqual(e.Param.String(), value), nil
	case *lep.GreaterThanX:
		return sql.GreaterThan(e.Param.String(), e.Value.Value()), nil
	case *lep.InSliceX:
		var items []interface{}
		for _, value := range e.Slice.Values {
			items = append(items, value.Value())
		}
		return sql.In(e.Param.String(), items...), nil

		// TODO: other cases
	}
}

func main() {
	query := `active=true && email!=null && (last_login>dt:"2010-01-01" || role in ["client","customer"])`

	expr, err := lep.ParseExpression(query)
	if err != nil {
		panic(err)
	}

	sql := sb.Select("*").From("users")
	where, err := traverse(sql, expr)
	if err != nil {
		panic(err)
	}
	sql.Where(where)

	spew.Dump(sql.Build())
}
```

```
(string) (len=99) "SELECT * FROM users WHERE (active = ? AND email IS NOT NULL AND (last_login > ? OR role IN (?, ?)))"
([]interface {}) (len=4 cap=4) {
 (bool) true,
 (time.Time) 2010-01-01 00:00:00 +0000 UTC,
 (string) (len=6) "client",
 (string) (len=8) "customer"
}
```
</details>

## Operators and types

* Comparators: `=` `!=` `>` `>=` `<` `<=` (left - param, right - param or value)
* Logical operations: `||` `&&` (left, right - any statements)
* Numeric constants: integer 64-bit (`12345678`), float 64-bit with floating point (`12345.678`)
* String constants (double quotes: `"foo bar"`, `"foo "bar""`)
* Regexp operations: `=~` (match regexp `a =~ /[a-z]+/`), `!~` (not match `b !~ /[0-9]+/`)
* Date constants (double quotes after `dt:`): `dt:"2020-03-04 10:20:30"` (for parsing datetime used [dateparse](github.com/araddon/dateparse) lib)
* Arrays (any values separated by `,` within square bracket: `[1,2,"foo",dt:"1999-09-09"]`)
* Array operations: `in` `not_in` (`a in [1,2,3]`)
* Boolean constants: `true` `false`
* Null constant: `null`

## Benchmarks

Here are the results output from a benchmark run on a Macbook Pro 2018:

```
go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/mgudov/logic-expression-parser
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkSmallQuery-12                     27211             42736 ns/op           20504 B/op        339 allocs/op
BenchmarkMediumQuery-12                    16160             72034 ns/op           32370 B/op        564 allocs/op
BenchmarkLargeQuery-12                      4066            286038 ns/op          114507 B/op       2134 allocs/op
BenchmarkSmallQueryWithMemo-12             13236             91061 ns/op           95044 B/op        283 allocs/op
BenchmarkMediumQueryWithMemo-12             4762            218919 ns/op          213942 B/op        626 allocs/op
BenchmarkLargeQueryWithMemo-12              1881            549937 ns/op          514177 B/op       1413 allocs/op
PASS
ok      github.com/mgudov/logic-expression-parser       9.347s
```

## Used Libraries

For parsing string the [pigeon](https://github.com/mna/pigeon) parser generator is used
(Licensed under [BSD 3-Clause](http://opensource.org/licenses/BSD-3-Clause)).
