package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGreaterThan(t *testing.T) {
	type testParseGreaterThan struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseGreaterThan{
		{
			left:   NewParam("a"),
			right:  NewInteger(100),
			result: `a>100`,
		},
		{
			left:   NewParam("a"),
			right:  NewInteger(-100),
			result: `a>-100`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(12.345),
			result: `a>12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(-12.345),
			result: `a>-12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewParam("b"),
			result: `a>b`,
		},
		{
			left:  NewInteger(1),
			right: NewParam("b"),
			err:   NewIncorrectType("parseStatement", (*Param)(nil), (*Integer)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseGreaterThan(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*GreaterThan)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.right, e.Value)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseGreaterThanEqual(t *testing.T) {
	type testParseGreaterThanEqual struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseGreaterThanEqual{
		{
			left:   NewParam("a"),
			right:  NewInteger(100),
			result: `a>=100`,
		},
		{
			left:   NewParam("a"),
			right:  NewInteger(-100),
			result: `a>=-100`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(12.345),
			result: `a>=12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(-12.345),
			result: `a>=-12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewParam("b"),
			result: `a>=b`,
		},
		{
			left:  NewInteger(1),
			right: NewParam("b"),
			err:   NewIncorrectType("parseStatement", (*Param)(nil), (*Integer)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseGreaterThanEqual(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*GreaterThanEqual)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.right, e.Value)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestGreaterThan_Equals(t *testing.T) {
	type testGreaterThanEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testGreaterThanEquals{
		{
			e1:     NewGreaterThan(NewParam("a"), NewInteger(1000)),
			e2:     NewGreaterThan(NewParam("a"), NewInteger(1000)),
			result: true,
		},
		{
			e1:     NewGreaterThan(NewParam("a"), NewFloat(12.345)),
			e2:     NewGreaterThan(NewParam("a"), NewFloat(12.345)),
			result: true,
		},
		{
			e1:     NewGreaterThan(NewParam("a"), NewInteger(1000)),
			e2:     NewGreaterThan(NewParam("a"), NewFloat(1000)),
			result: false,
		},
		{
			e1:     NewGreaterThan(NewParam("a"), NewInteger(1000)),
			e2:     NewEquals(NewParam("a"), NewInteger(1000)),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}

func TestGreaterThanEqual_Equals(t *testing.T) {
	type testGreaterThanEqualEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testGreaterThanEqualEquals{
		{
			e1:     NewGreaterThanEqual(NewParam("a"), NewInteger(1000)),
			e2:     NewGreaterThanEqual(NewParam("a"), NewInteger(1000)),
			result: true,
		},
		{
			e1:     NewGreaterThanEqual(NewParam("a"), NewFloat(12.345)),
			e2:     NewGreaterThanEqual(NewParam("a"), NewFloat(12.345)),
			result: true,
		},
		{
			e1:     NewGreaterThanEqual(NewParam("a"), NewInteger(1000)),
			e2:     NewGreaterThanEqual(NewParam("a"), NewFloat(1000)),
			result: false,
		},
		{
			e1:     NewGreaterThanEqual(NewParam("a"), NewInteger(1000)),
			e2:     NewEquals(NewParam("a"), NewInteger(1000)),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}
