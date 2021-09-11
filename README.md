logic-expression-parser
====

[![Build Status](https://app.travis-ci.com/mgudov/logic-expression-parser.svg?branch=master)](https://app.travis-ci.com/mgudov/logic-expression-parser)
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
(*lep.And)({
 conjuncts: ([]lep.Expression) (len=3 cap=4) {
  (*lep.Equals)({
   Param: (*lep.Param)({
    Name: (string) (len=1) "a"
   }),
   Value: (*lep.Boolean)({
    Val: (bool) false
   })
  }),
  (*lep.GreaterThanEqual)({
   Param: (*lep.Param)({
    Name: (string) (len=1) "b"
   }),
   Value: (*lep.Param)({
    Name: (string) (len=1) "c"
   })
  }),
  (*lep.Or)({
   disjunctions: ([]lep.Expression) (len=2 cap=2) {
    (*lep.LessThan)({
     Param: (*lep.Param)({
      Name: (string) (len=1) "d"
     }),
     Value: (*lep.Integer)({
      Val: (int64) 1000
     })
    }),
    (*lep.InSlice)({
     Param: (*lep.Param)({
      Name: (string) (len=1) "e"
     }),
     Slice: (*lep.Slice)({
      Values: ([]lep.Value) (len=3 cap=4) {
       (*lep.Integer)({
        Val: (int64) 1
       }),
       (*lep.Integer)({
        Val: (int64) 2
       }),
       (*lep.Integer)({
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
	expression := lep.NewAnd(
		lep.NewEquals(lep.NewParam("a"), lep.NewBoolean(false)),
		lep.NewGreaterThanEqual(lep.NewParam("b"), lep.NewParam("c")),
		lep.NewOr(
			lep.NewLessThan(lep.NewParam("d"), lep.NewInteger(1000)),
			lep.NewInSlice(
				lep.NewParam("e"),
				lep.NewSlice(lep.NewInteger(1), lep.NewInteger(2), lep.NewInteger(3)),
			),
		),
	)

	fmt.Println(expression.String())
}
```

```
a=false && b>=c && (d<1000 || e in [1,2,3])
```

## Benchmarks

Here are the results output from a benchmark run on a Macbook Pro 2018:

```
go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/mgudov/logic-expression-parser
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkSmallQuery-12             	   22050	     57427 ns/op	   20600 B/op	     343 allocs/op
BenchmarkMediumQuery-12            	   16213	     77054 ns/op	   32560 B/op	     572 allocs/op
BenchmarkLargeQuery-12             	    4184	    320364 ns/op	  114753 B/op	    2144 allocs/op
BenchmarkSmallQueryWithMemo-12     	   10000	    100139 ns/op	   95046 B/op	     284 allocs/op
BenchmarkMediumQueryWithMemo-12    	    5685	    233060 ns/op	  214027 B/op	     632 allocs/op
BenchmarkLargeQueryWithMemo-12     	    2136	    561411 ns/op	  514323 B/op	    1417 allocs/op
PASS
ok  	github.com/mgudov/logic-expression-parser	8.928s
```

## Used Libraries

For parsing the string query binding the [pigeon](https://github.com/mna/pigeon) parser generator is used
(Licensed under [BSD 3-Clause](http://opensource.org/licenses/BSD-3-Clause)).
