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
			left:   Param("a"),
			right:  Integer(100),
			result: `a<100`,
		},
		{
			left:   Param("a"),
			right:  Integer(-100),
			result: `a<-100`,
		},
		{
			left:   Param("a"),
			right:  Float(12.345),
			result: `a<12.345`,
		},
		{
			left:   Param("a"),
			right:  Float(-12.345),
			result: `a<-12.345`,
		},
		{
			left:   Param("a"),
			right:  Param("b"),
			result: `a<b`,
		},
		{
			left:  Integer(1),
			right: Param("b"),
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseLessThan(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*LessThanX)(nil), e)
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
			left:   Param("a"),
			right:  Integer(100),
			result: `a<=100`,
		},
		{
			left:   Param("a"),
			right:  Integer(-100),
			result: `a<=-100`,
		},
		{
			left:   Param("a"),
			right:  Float(12.345),
			result: `a<=12.345`,
		},
		{
			left:   Param("a"),
			right:  Float(-12.345),
			result: `a<=-12.345`,
		},
		{
			left:   Param("a"),
			right:  Param("b"),
			result: `a<=b`,
		},
		{
			left:  Integer(1),
			right: Param("b"),
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseLessThanEqual(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*LessThanEqualX)(nil), e)
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
			e1:     LessThan(Param("a"), Integer(1000)),
			e2:     LessThan(Param("a"), Integer(1000)),
			result: true,
		},
		{
			e1:     LessThan(Param("a"), Float(12.345)),
			e2:     LessThan(Param("a"), Float(12.345)),
			result: true,
		},
		{
			e1:     LessThan(Param("a"), Integer(1000)),
			e2:     LessThan(Param("a"), Float(1000)),
			result: false,
		},
		{
			e1:     LessThan(Param("a"), Integer(1000)),
			e2:     Equals(Param("a"), Integer(1000)),
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
			e1:     LessThanEqual(Param("a"), Integer(1000)),
			e2:     LessThanEqual(Param("a"), Integer(1000)),
			result: true,
		},
		{
			e1:     LessThanEqual(Param("a"), Float(12.345)),
			e2:     LessThanEqual(Param("a"), Float(12.345)),
			result: true,
		},
		{
			e1:     LessThanEqual(Param("a"), Integer(1000)),
			e2:     LessThanEqual(Param("a"), Float(1000)),
			result: false,
		},
		{
			e1:     LessThanEqual(Param("a"), Integer(1000)),
			e2:     Equals(Param("a"), Integer(1000)),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}
