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
			left:   Param("a"),
			right:  String("foo"),
			result: `a="foo"`,
		},
		{
			left:   Param("a"),
			right:  Integer(1000),
			result: `a=1000`,
		},
		{
			left:   Param("a"),
			right:  Float(12.345),
			result: `a=12.345`,
		},
		{
			left:   Param("a"),
			right:  Param("b"),
			result: `a=b`,
		},
		{
			left:  Integer(1),
			right: Param("b"),
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseEquals(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*EqualsX)(nil), e)
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
			left:   Param("a"),
			right:  String("foo"),
			result: `a!="foo"`,
		},
		{
			left:   Param("a"),
			right:  Integer(1000),
			result: `a!=1000`,
		},
		{
			left:   Param("a"),
			right:  Float(12.345),
			result: `a!=12.345`,
		},
		{
			left:   Param("a"),
			right:  Param("b"),
			result: `a!=b`,
		},
		{
			left:  Integer(1),
			right: Param("b"),
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseNotEquals(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && e != nil {
			assert.IsType(t, (*NotEqualsX)(nil), e)
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
			e1:     Equals(Param("a"), String("foo")),
			e2:     Equals(Param("a"), String("foo")),
			result: true,
		},
		{
			e1:     Equals(Param("a"), Integer(100)),
			e2:     Equals(Param("a"), Integer(100)),
			result: true,
		},
		{
			e1:     Equals(Param("a"), String("foo")),
			e2:     Equals(Param("a"), String("bar")),
			result: false,
		},
		{
			e1:     Equals(Param("a"), Integer(100)),
			e2:     Equals(Param("a"), Float(100)),
			result: false,
		},
		{
			e1:     Equals(Param("a"), Integer(100)),
			e2:     Integer(100),
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
			e1:     NotEquals(Param("a"), String("foo")),
			e2:     NotEquals(Param("a"), String("foo")),
			result: true,
		},
		{
			e1:     NotEquals(Param("a"), Integer(100)),
			e2:     NotEquals(Param("a"), Integer(100)),
			result: true,
		},
		{
			e1:     NotEquals(Param("a"), String("foo")),
			e2:     NotEquals(Param("a"), String("bar")),
			result: false,
		},
		{
			e1:     NotEquals(Param("a"), Integer(100)),
			e2:     NotEquals(Param("a"), Float(100)),
			result: false,
		},
		{
			e1:     NotEquals(Param("a"), Integer(100)),
			e2:     Integer(100),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}
