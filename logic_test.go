package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseExpressions(t *testing.T) {
	var (
		e1 = NewEquals(NewParam("a"), NewString("foo"))
		e2 = NewGreaterThan(NewParam("b"), NewInteger(100))
		e3 = NewLessThan(NewParam("c"), NewFloat(12.34))
		e4 = NewNotEquals(NewParam("d"), NewBoolean(true))
		e5 = NewNotEquals(NewParam("e"), NewNull())
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
		e1 = NewEquals(NewParam("a"), NewString("foo"))
		e2 = NewGreaterThan(NewParam("b"), NewInteger(100))
		e3 = NewLessThan(NewParam("c"), NewFloat(12.34))
		e4 = NewNotEquals(NewParam("d"), NewBoolean(true))
		e5 = NewNotEquals(NewParam("e"), NewNull())
	)

	type testParseAnd struct {
		expr   []interface{}
		result Expression
	}
	var tests = []testParseAnd{
		{
			expr:   []interface{}{e1, e2, e3, e4, e5},
			result: NewAnd(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{NewAnd(e1, e2), NewAnd(e3, e4, e5)},
			result: NewAnd(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{NewAnd(e1, e2), NewAnd(e3, e4, e5)},
			result: NewAnd(NewAnd(e1, e2), NewAnd(e3, e4, e5)),
		},
		{
			expr:   []interface{}{NewAnd(e1, e2), NewAnd(e3, e4, e5)},
			result: NewAnd(NewAnd(e1, e2, e3), NewAnd(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, NewOr(e4, e5)},
			result: NewAnd(e1, e2, e3, NewOr(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, NewOr(e4, e5)},
			result: NewAnd(e1, e2, NewAnd(e3, NewOr(e4, e5))),
		},
	}

	for _, tt := range tests {
		l, err := parseAnd(tt.expr...)
		if assert.NoError(t, err) {
			assert.Equal(t, tt.result, l)
			assert.IsType(t, (*And)(nil), l)
			assert.Equal(t, tt.result.String(), l.String())
		}
	}
}

func TestParseOr(t *testing.T) {
	var (
		e1 = NewEquals(NewParam("a"), NewString("foo"))
		e2 = NewGreaterThan(NewParam("b"), NewInteger(100))
		e3 = NewLessThan(NewParam("c"), NewFloat(12.34))
		e4 = NewNotEquals(NewParam("d"), NewBoolean(true))
		e5 = NewNotEquals(NewParam("e"), NewNull())
	)

	type testParseOr struct {
		expr   []interface{}
		result Expression
	}
	var tests = []testParseOr{
		{
			expr:   []interface{}{e1, e2, e3, e4, e5},
			result: NewOr(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{NewOr(e1, e2), NewOr(e3, e4, e5)},
			result: NewOr(e1, e2, e3, e4, e5),
		},
		{
			expr:   []interface{}{NewOr(e1, e2), NewOr(e3, e4, e5)},
			result: NewOr(NewOr(e1, e2), NewOr(e3, e4, e5)),
		},
		{
			expr:   []interface{}{NewOr(e1, e2), NewOr(e3, e4, e5)},
			result: NewOr(NewOr(e1, e2, e3), NewOr(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, NewAnd(e4, e5)},
			result: NewOr(e1, e2, e3, NewAnd(e4, e5)),
		},
		{
			expr:   []interface{}{e1, e2, e3, NewAnd(e4, e5)},
			result: NewOr(e1, e2, NewOr(e3, NewAnd(e4, e5))),
		},
	}

	for _, tt := range tests {
		l, err := parseOr(tt.expr...)
		if assert.NoError(t, err) {
			assert.Equal(t, tt.result, l)
			assert.IsType(t, (*Or)(nil), l)
			assert.Equal(t, l.String(), tt.result.String())
		}
	}
}

func TestAnd_Equals(t *testing.T) {
	var (
		e1 = NewEquals(NewParam("a"), NewString("foo"))
		e2 = NewGreaterThan(NewParam("b"), NewInteger(100))
		e3 = NewLessThan(NewParam("c"), NewFloat(12.34))
		e4 = NewNotEquals(NewParam("d"), NewBoolean(true))
		e5 = NewNotEquals(NewParam("e"), NewNull())
	)

	type testAndEquals struct {
		l1     Expression
		l2     Expression
		result bool
	}
	var tests = []testAndEquals{
		{
			l1:     NewAnd(e1, e2, e3, e4, e5),
			l2:     NewAnd(e1, e2, e3, e4, e5),
			result: true,
		},
		{
			l1:     NewAnd(e1, e2, e3, e4, e5),
			l2:     NewAnd(NewAnd(e1, e2), NewAnd(e3, e4), e5),
			result: true,
		},
		{
			l1:     NewAnd(e1, e2, e3, e4, e5),
			l2:     NewAnd(e1, e2, e3, e4),
			result: false,
		},
		{
			l1:     NewAnd(e1, e2, e3, e4, e5),
			l2:     NewAnd(NewAnd(e1, e2), NewAnd(e3, e4)),
			result: false,
		},
		{
			l1:     NewAnd(e1, e2, e3, e4, e5),
			l2:     NewOr(e1, e2, e3, e4, e5),
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
		e1 = NewEquals(NewParam("a"), NewString("foo"))
		e2 = NewGreaterThan(NewParam("b"), NewInteger(100))
		e3 = NewLessThan(NewParam("c"), NewFloat(12.34))
		e4 = NewNotEquals(NewParam("d"), NewBoolean(true))
		e5 = NewNotEquals(NewParam("e"), NewNull())
	)

	type testOrEquals struct {
		l1     Expression
		l2     Expression
		result bool
	}
	var tests = []testOrEquals{
		{
			l1:     NewOr(e1, e2, e3, e4, e5),
			l2:     NewOr(e1, e2, e3, e4, e5),
			result: true,
		},
		{
			l1:     NewOr(e1, e2, e3, e4, e5),
			l2:     NewOr(NewOr(e1, e2), NewOr(e3, e4), e5),
			result: true,
		},
		{
			l1:     NewOr(e1, e2, e3, e4, e5),
			l2:     NewOr(e1, e2, e3, e4),
			result: false,
		},
		{
			l1:     NewOr(e1, e2, e3, e4, e5),
			l2:     NewOr(NewOr(e1, e2), NewOr(e3, e4)),
			result: false,
		},
		{
			l1:     NewOr(e1, e2, e3, e4, e5),
			l2:     NewAnd(e1, e2, e3, e4, e5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.l1.Equals(tt.l2))
		assert.Equal(t, tt.result, tt.l2.Equals(tt.l1))
	}
}
