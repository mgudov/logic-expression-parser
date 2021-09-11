package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLessThan(t *testing.T) {
	type testParseLessThan struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseLessThan{
		{
			left:   NewParam("a"),
			right:  NewInteger(100),
			result: `a<100`,
		},
		{
			left:   NewParam("a"),
			right:  NewInteger(-100),
			result: `a<-100`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(12.345),
			result: `a<12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(-12.345),
			result: `a<-12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewParam("b"),
			result: `a<b`,
		},
		{
			left:  NewInteger(1),
			right: NewParam("b"),
			err:   NewIncorrectType("parseStatement", (*Param)(nil), (*Integer)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseLessThan(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*LessThan)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.right, e.Value)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseLessThanEqual(t *testing.T) {
	type testParseLessThanEqual struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseLessThanEqual{
		{
			left:   NewParam("a"),
			right:  NewInteger(100),
			result: `a<=100`,
		},
		{
			left:   NewParam("a"),
			right:  NewInteger(-100),
			result: `a<=-100`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(12.345),
			result: `a<=12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewFloat(-12.345),
			result: `a<=-12.345`,
		},
		{
			left:   NewParam("a"),
			right:  NewParam("b"),
			result: `a<=b`,
		},
		{
			left:  NewInteger(1),
			right: NewParam("b"),
			err:   NewIncorrectType("parseStatement", (*Param)(nil), (*Integer)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseLessThanEqual(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*LessThanEqual)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.right, e.Value)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestLessThan_Equals(t *testing.T) {
	type testLessThanEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testLessThanEquals{
		{
			e1:     NewLessThan(NewParam("a"), NewInteger(1000)),
			e2:     NewLessThan(NewParam("a"), NewInteger(1000)),
			result: true,
		},
		{
			e1:     NewLessThan(NewParam("a"), NewFloat(12.345)),
			e2:     NewLessThan(NewParam("a"), NewFloat(12.345)),
			result: true,
		},
		{
			e1:     NewLessThan(NewParam("a"), NewInteger(1000)),
			e2:     NewLessThan(NewParam("a"), NewFloat(1000)),
			result: false,
		},
		{
			e1:     NewLessThan(NewParam("a"), NewInteger(1000)),
			e2:     NewEquals(NewParam("a"), NewInteger(1000)),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}

func TestLessThanEqual_Equals(t *testing.T) {
	type testGreaterThanEqualEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testGreaterThanEqualEquals{
		{
			e1:     NewLessThanEqual(NewParam("a"), NewInteger(1000)),
			e2:     NewLessThanEqual(NewParam("a"), NewInteger(1000)),
			result: true,
		},
		{
			e1:     NewLessThanEqual(NewParam("a"), NewFloat(12.345)),
			e2:     NewLessThanEqual(NewParam("a"), NewFloat(12.345)),
			result: true,
		},
		{
			e1:     NewLessThanEqual(NewParam("a"), NewInteger(1000)),
			e2:     NewLessThanEqual(NewParam("a"), NewFloat(1000)),
			result: false,
		},
		{
			e1:     NewLessThanEqual(NewParam("a"), NewInteger(1000)),
			e2:     NewEquals(NewParam("a"), NewInteger(1000)),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}
