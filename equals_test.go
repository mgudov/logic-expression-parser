package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEquals(t *testing.T) {
	type testParseEquals struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseEquals{
		{
			left:   NewParam("a"),
			right:  NewString("foo"),
			result: `a="foo"`,
		},
		{
			left:   NewParam("a"),
			right:  NewInteger(1000),
			result: `a=1000`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(12.345),
			result: `a=12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewParam("b"),
			result: `a=b`,
		},
		{
			left:  NewInteger(1),
			right: NewParam("b"),
			err:   NewIncorrectType("parseStatement", (*Param)(nil), (*Integer)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseEquals(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*Equals)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.right, e.Value)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseNotEquals(t *testing.T) {
	type testParseNotEquals struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseNotEquals{
		{
			left:   NewParam("a"),
			right:  NewString("foo"),
			result: `a!="foo"`,
		},
		{
			left:   NewParam("a"),
			right:  NewInteger(1000),
			result: `a!=1000`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(12.345),
			result: `a!=12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewParam("b"),
			result: `a!=b`,
		},
		{
			left:  NewInteger(1),
			right: NewParam("b"),
			err:   NewIncorrectType("parseStatement", (*Param)(nil), (*Integer)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseNotEquals(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*NotEquals)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.right, e.Value)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestEquals_Equals(t *testing.T) {
	type testEqualsEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testEqualsEquals{
		{
			e1:     NewEquals(NewParam("a"), NewString("foo")),
			e2:     NewEquals(NewParam("a"), NewString("foo")),
			result: true,
		},
		{
			e1:     NewEquals(NewParam("a"), NewInteger(100)),
			e2:     NewEquals(NewParam("a"), NewInteger(100)),
			result: true,
		},
		{
			e1:     NewEquals(NewParam("a"), NewString("foo")),
			e2:     NewEquals(NewParam("a"), NewString("bar")),
			result: false,
		},
		{
			e1:     NewEquals(NewParam("a"), NewInteger(100)),
			e2:     NewEquals(NewParam("a"), NewFloat(100)),
			result: false,
		},
		{
			e1:     NewEquals(NewParam("a"), NewInteger(100)),
			e2:     NewInteger(100),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}

func TestNotEquals_Equals(t *testing.T) {
	type testNotEqualsEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testNotEqualsEquals{
		{
			e1:     NewNotEquals(NewParam("a"), NewString("foo")),
			e2:     NewNotEquals(NewParam("a"), NewString("foo")),
			result: true,
		},
		{
			e1:     NewNotEquals(NewParam("a"), NewInteger(100)),
			e2:     NewNotEquals(NewParam("a"), NewInteger(100)),
			result: true,
		},
		{
			e1:     NewNotEquals(NewParam("a"), NewString("foo")),
			e2:     NewNotEquals(NewParam("a"), NewString("bar")),
			result: false,
		},
		{
			e1:     NewNotEquals(NewParam("a"), NewInteger(100)),
			e2:     NewNotEquals(NewParam("a"), NewFloat(100)),
			result: false,
		},
		{
			e1:     NewNotEquals(NewParam("a"), NewInteger(100)),
			e2:     NewInteger(100),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}
