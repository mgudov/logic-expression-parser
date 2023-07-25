package lep

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestParseExpression(t *testing.T) {
	var (
		a = Param("a")
		b = Param("b")
		c = Param("c")
		d = Param("d")
		g = Param("g")
		e = Param("e")
		f = Param("f")
	)

	type testParseExpression struct {
		query string
		expr  Expression
		err   error
	}
	var tests = []testParseExpression{
		{
			query: `a=10`,
			expr:  Equals(a, Integer(10)),
		},
		{
			query: `a="Test"`,
			expr:  Equals(a, String("Test")),
		},
		{
			query: `a>100 && a<200`,
			expr: And(
				GreaterThan(a, Integer(100)),
				LessThan(a, Integer(200)),
			),
		},
		{
			query: `a>=100 && a<=200 || b="Foo" && c!="bar"`,
			expr: Or(
				And(
					GreaterThanEqual(a, Integer(100)),
					LessThanEqual(a, Integer(200)),
				),
				And(
					Equals(b, String("Foo")),
					NotEquals(c, String("bar")),
				),
			),
		},
		{
			query: `a in [100,10.5,"some text"] && b not_in [false,null,12.34,-56.78]`,
			expr: And(
				InSlice(a, Slice(Integer(100), Float(10.5), String("some text"))),
				NotInSlice(b, Slice(Boolean(false), Null(), Float(12.34), Float(-56.78))),
			),
		},
		{
			query: `a has 100 && b not_has null || a has_any ["aaa","bbb","ccc"] && b has_all [111,222,333]`,
			expr: Or(
				And(Has(a, Integer(100)), NotHas(b, Null())),
				And(
					HasAny(a, Slice(String("aaa"), String("bbb"), String("ccc"))),
					HasAll(b, Slice(Integer(111), Integer(222), Integer(333))),
				),
			),
		},
		{
			query: `(a!=null || b=false) && (b!=true || c=null) && (a=null || b=true || c=false)`,
			expr: And(
				Or(NotEquals(a, Null()), Equals(b, Boolean(false))),
				Or(NotEquals(b, Boolean(true)), Equals(c, Null())),
				Or(Equals(a, Null()), Equals(b, Boolean(true)), Equals(c, Boolean(false))),
			),
		},
		{
			query: `a<dt:"2021-05-01" || b>dt:"2021-01-09" || c=dt:"1988-10-30"`,
			expr: Or(
				LessThan(a, DateTime(time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC), "2006-01-02")),
				GreaterThan(b, DateTime(time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC), "2006-01-02")),
				Equals(c, DateTime(time.Date(1988, 10, 30, 0, 0, 0, 0, time.UTC), "2006-01-02")),
			),
		},
		{
			query: `(a!=b || a!=null || a!="" || a!=0) && (b=a || b=null || b="" || b=0)`,
			expr: And(
				Or(NotEquals(a, b), NotEquals(a, Null()), NotEquals(a, String("")), NotEquals(a, Integer(0))),
				Or(Equals(b, a), Equals(b, Null()), Equals(b, String("")), Equals(b, Integer(0))),
			),
		},
		{
			query: `b>=c && (a=123 || d=345 || g!=null) || e<123 && f in [1,2,3,4,5]`,
			expr: Or(
				And(
					GreaterThanEqual(b, c),
					Or(Equals(a, Integer(123)), Equals(d, Integer(345)), NotEquals(g, Null())),
				),
				And(
					LessThan(e, Integer(123)),
					InSlice(f, Slice(Integer(1), Integer(2), Integer(3), Integer(4), Integer(5))),
				),
			),
		},
		{
			query: `a =~ /[a-zA-Z]+/ && b !~ /[0-9]+/gm && (c =~ /[\d]+/ || d !~ /[\D]+/)`,
			expr: And(
				MatchRegexp(a, Regexp(regexp.MustCompile(`/[a-zA-Z]+/`))),
				NotMatchRegexp(b, Regexp(regexp.MustCompile(`/[0-9]+/gm`))),
				Or(
					MatchRegexp(c, Regexp(regexp.MustCompile(`/[\d]+/`))),
					NotMatchRegexp(d, Regexp(regexp.MustCompile(`/[\D]+/`))),
				),
			),
		},
		{
			query: `a=="undefined operator"`,
			err:   errors.New("no match found"),
		},
	}

	for _, tt := range tests {
		expr, err := ParseExpression(tt.query)
		if tt.err == nil && assert.NoError(t, err) {
			assert.Equal(t, tt.expr, expr)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestExpression_String(t *testing.T) {
	type testExpressionString struct {
		query  string
		result string
	}
	var tests = []testExpressionString{
		{
			query:  `a >= 100 && a <= 200 || b = "Foo" && c != "bar"`,
			result: `a>=100 && a<=200 || b="Foo" && c!="bar"`,
		},
		{
			query:  `a in [100,10.5,"some text"] && b not_in [false,null,12.34,-56.78]`,
			result: `a in [100,10.5,"some text"] && b not_in [false,null,12.34,-56.78]`,
		},
		{
			query:  `(a != null || (b=false)) && (((b!=true)) || c=null) && (a = null || b = true || c = false)`,
			result: `(a!=null || b=false) && (b!=true || c=null) && (a=null || b=true || c=false)`,
		},
		{
			query:  `(a<dt:"2021-05-01" || b > dt:"2021-01-09") || c = dt:"1988-10-30"`,
			result: `a<dt:"2021-05-01" || b>dt:"2021-01-09" || c=dt:"1988-10-30"`,
		},
		{
			query:  `((b >= c) && (a=123 || g=345 || j!=null)) || t<123 && k in [1,2,3,4,5]`,
			result: `b>=c && (a=123 || g=345 || j!=null) || t<123 && k in [1,2,3,4,5]`,
		},
	}

	for _, tt := range tests {
		expr, err := ParseExpression(tt.query)
		if assert.NoError(t, err) {
			assert.Equal(t, tt.result, expr.String())
		}
	}
}
