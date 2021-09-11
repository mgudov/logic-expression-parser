package lep

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseExpression(t *testing.T) {
	var (
		a = NewParam("a")
		b = NewParam("b")
		c = NewParam("c")
		d = NewParam("d")
		g = NewParam("g")
		e = NewParam("e")
		f = NewParam("f")
	)

	type testParseExpression struct {
		query string
		expr  Expression
		err   error
	}
	var tests = []testParseExpression{
		{
			query: `a=10`,
			expr:  NewEquals(a, NewInteger(10)),
		},
		{
			query: `a="Test"`,
			expr:  NewEquals(a, NewString("Test")),
		},
		{
			query: `a>100 && a<200`,
			expr: NewAnd(
				NewGreaterThan(a, NewInteger(100)),
				NewLessThan(a, NewInteger(200)),
			),
		},
		{
			query: `a>=100 && a<=200 || b="Foo" && c!="bar"`,
			expr: NewOr(
				NewAnd(
					NewGreaterThanEqual(a, NewInteger(100)),
					NewLessThanEqual(a, NewInteger(200)),
				),
				NewAnd(
					NewEquals(b, NewString("Foo")),
					NewNotEquals(c, NewString("bar")),
				),
			),
		},
		{
			query: `a in [100,10.5,"some text"] && b not_in [false,null,12.34,-56.78]`,
			expr: NewAnd(
				NewInSlice(a, NewSlice(NewInteger(100), NewFloat(10.5), NewString("some text"))),
				NewNotInSlice(b, NewSlice(NewBoolean(false), NewNull(), NewFloat(12.34), NewFloat(-56.78))),
			),
		},
		{
			query: `(a!=null || b=false) && (b!=true || c=null) && (a=null || b=true || c=false)`,
			expr: NewAnd(
				NewOr(NewNotEquals(a, NewNull()), NewEquals(b, NewBoolean(false))),
				NewOr(NewNotEquals(b, NewBoolean(true)), NewEquals(c, NewNull())),
				NewOr(NewEquals(a, NewNull()), NewEquals(b, NewBoolean(true)), NewEquals(c, NewBoolean(false))),
			),
		},
		{
			query: `a<dt:"2021-05-01" || b>dt:"2021-01-09" || c=dt:"1988-10-30"`,
			expr: NewOr(
				NewLessThan(a, NewDateTime(time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC), "2006-01-02")),
				NewGreaterThan(b, NewDateTime(time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC), "2006-01-02")),
				NewEquals(c, NewDateTime(time.Date(1988, 10, 30, 0, 0, 0, 0, time.UTC), "2006-01-02")),
			),
		},
		{
			query: `b>=c && (a=123 || d=345 || g!=null) || e<123 && f in [1,2,3,4,5]`,
			expr: NewOr(
				NewAnd(
					NewGreaterThanEqual(b, c),
					NewOr(NewEquals(a, NewInteger(123)), NewEquals(d, NewInteger(345)), NewNotEquals(g, NewNull())),
				),
				NewAnd(
					NewLessThan(e, NewInteger(123)),
					NewInSlice(f, NewSlice(NewInteger(1), NewInteger(2), NewInteger(3), NewInteger(4), NewInteger(5))),
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
