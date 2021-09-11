package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseExpressions(t *testing.T) {
	var (
		e1 = Equals(Param("a"), String("foo"))
		e2 = GreaterThan(Param("b"), Integer(100))
		e3 = LessThan(Param("c"), Float(12.34))
		e4 = NotEquals(Param("d"), Boolean(true))
		e5 = NotEquals(Param("e"), Null())
	)

	type testParseExpressions struct {
		items []interface{}
		expr  []Expression
	}
	var tests = []testParseExpressions{
		{
			items: []interface{}{e1, e2, e3, e4, e5},
			expr:  []Expression{e1, e2, e3, e4, e5},
		},
		{
			items: []interface{}{e1, []interface{}{e2, e3}, []interface{}{e4, e5}},
			expr:  []Expression{e1, e2, e3, e4, e5},
		},
		{
			items: []interface{}{"foo", e1, "bar"},
			expr:  []Expression{e1},
		},
		{
			items: []interface{}{"foo", "bar"},
			expr:  nil,
		},
	}

	for _, tt := range tests {
		expr := parseExpressions(tt.items...)
		assert.Equal(t, tt.expr, expr)
	}
}

func TestParseAnd(t *testing.T) {
	var (
		e1 = Equals(Param("a"), String("foo"))
		e2 = GreaterThan(Param("b"), Integer(100))
		e3 = LessThan(Param("c"), Float(12.34))
		e4 = NotEquals(Param("d"), Boolean(true))
		e5 = NotEquals(Param("e"), Null())
	)

	type testParseAnd struct {
		expr   []interface{}
		result Expression
	}
	var tests = []testParseAnd{
		{
			expr:   []interface{}{e1, e2, e3, e4, e5},
			result: And(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{And(e1, e2), And(e3, e4, e5)},
			result: And(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{And(e1, e2), And(e3, e4, e5)},
			result: And(And(e1, e2), And(e3, e4, e5)),
		},
		{
			expr:   []interface{}{And(e1, e2), And(e3, e4, e5)},
			result: And(And(e1, e2, e3), And(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, Or(e4, e5)},
			result: And(e1, e2, e3, Or(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, Or(e4, e5)},
			result: And(e1, e2, And(e3, Or(e4, e5))),
		},
	}

	for _, tt := range tests {
		l, err := parseAnd(tt.expr...)
		if assert.NoError(t, err) {
			assert.Equal(t, tt.result, l)
			assert.IsType(t, (*AndX)(nil), l)
			assert.Equal(t, tt.result.String(), l.String())
		}
	}
}

func TestParseOr(t *testing.T) {
	var (
		e1 = Equals(Param("a"), String("foo"))
		e2 = GreaterThan(Param("b"), Integer(100))
		e3 = LessThan(Param("c"), Float(12.34))
		e4 = NotEquals(Param("d"), Boolean(true))
		e5 = NotEquals(Param("e"), Null())
	)

	type testParseOr struct {
		expr   []interface{}
		result Expression
	}
	var tests = []testParseOr{
		{
			expr:   []interface{}{e1, e2, e3, e4, e5},
			result: Or(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{Or(e1, e2), Or(e3, e4, e5)},
			result: Or(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{Or(e1, e2), Or(e3, e4, e5)},
			result: Or(Or(e1, e2), Or(e3, e4, e5)),
		},
		{
			expr:   []interface{}{Or(e1, e2), Or(e3, e4, e5)},
			result: Or(Or(e1, e2, e3), Or(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, And(e4, e5)},
			result: Or(e1, e2, e3, And(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, And(e4, e5)},
			result: Or(e1, e2, Or(e3, And(e4, e5))),
		},
	}

	for _, tt := range tests {
		l, err := parseOr(tt.expr...)
		if assert.NoError(t, err) {
			assert.Equal(t, tt.result, l)
			assert.IsType(t, (*OrX)(nil), l)
			assert.Equal(t, l.String(), tt.result.String())
		}
	}
}

func TestAnd_Equals(t *testing.T) {
	var (
		e1 = Equals(Param("a"), String("foo"))
		e2 = GreaterThan(Param("b"), Integer(100))
		e3 = LessThan(Param("c"), Float(12.34))
		e4 = NotEquals(Param("d"), Boolean(true))
		e5 = NotEquals(Param("e"), Null())
	)

	type testAndEquals struct {
		l1     Expression
		l2     Expression
		result bool
	}
	var tests = []testAndEquals{
		{
			l1:     And(e1, e2, e3, e4, e5),
			l2:     And(e1, e2, e3, e4, e5),
			result: true,
		},
		{
			l1:     And(e1, e2, e3, e4, e5),
			l2:     And(And(e1, e2), And(e3, e4), e5),
			result: true,
		},
		{
			l1:     And(e1, e2, e3, e4, e5),
			l2:     And(e1, e2, e3, e4),
			result: false,
		},
		{
			l1:     And(e1, e2, e3, e4, e5),
			l2:     And(And(e1, e2), And(e3, e4)),
			result: false,
		},
		{
			l1:     And(e1, e2, e3, e4, e5),
			l2:     Or(e1, e2, e3, e4, e5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.l1.Equals(tt.l2))
		assert.Equal(t, tt.result, tt.l2.Equals(tt.l1))
	}
}

func TestOr_Equals(t *testing.T) {
	var (
		e1 = Equals(Param("a"), String("foo"))
		e2 = GreaterThan(Param("b"), Integer(100))
		e3 = LessThan(Param("c"), Float(12.34))
		e4 = NotEquals(Param("d"), Boolean(true))
		e5 = NotEquals(Param("e"), Null())
	)

	type testOrEquals struct {
		l1     Expression
		l2     Expression
		result bool
	}
	var tests = []testOrEquals{
		{
			l1:     Or(e1, e2, e3, e4, e5),
			l2:     Or(e1, e2, e3, e4, e5),
			result: true,
		},
		{
			l1:     Or(e1, e2, e3, e4, e5),
			l2:     Or(Or(e1, e2), Or(e3, e4), e5),
			result: true,
		},
		{
			l1:     Or(e1, e2, e3, e4, e5),
			l2:     Or(e1, e2, e3, e4),
			result: false,
		},
		{
			l1:     Or(e1, e2, e3, e4, e5),
			l2:     Or(Or(e1, e2), Or(e3, e4)),
			result: false,
		},
		{
			l1:     Or(e1, e2, e3, e4, e5),
			l2:     And(e1, e2, e3, e4, e5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.l1.Equals(tt.l2))
		assert.Equal(t, tt.result, tt.l2.Equals(tt.l1))
	}
}
