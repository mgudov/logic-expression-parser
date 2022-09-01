package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArrayHasValue(t *testing.T) {
	type testParseArrayHasValue struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseArrayHasValue{
		{
			left:   Param("a"),
			right:  String("foo"),
			result: `a has "foo"`,
		},
		{
			left:   Param("a"),
			right:  Integer(1000),
			result: `a has 1000`,
		},
		{
			left:   Param("a"),
			right:  Float(12.345),
			result: `a has 12.345`,
		},
		{
			left:   Param("a"),
			right:  Param("b"),
			result: `a has b`,
		},
		{
			left:  Integer(1),
			right: Param("b"),
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseArrayHasValue(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*ArrayHasValueX)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.right, e.Value)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestArrayHasAllValuesParse(t *testing.T) {
	var (
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)
	type testParseArrayHasAllValues struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}

	var tests = []testParseArrayHasAllValues{
		{
			left:   Param("a"),
			right:  Slice(v1, v2, v3, v4, v5),
			result: `a has_all ["foo",100,12.34,true,null]`,
		},
		{
			left:   Param("a"),
			right:  Slice(v5),
			result: `a has_all [null]`,
		},
		{
			left:  Param("a"),
			right: v1,
			err:   IncorrectType("parseArrayHasAllValues", (*SliceX)(nil), (*StringX)(nil)),
		},
		{
			left:  Param("a"),
			right: NotEquals(Param("b"), v1),
			err:   IncorrectType("parseStatement", (*Value)(nil), (*NotEqualsX)(nil)),
		},
	}
	for _, tt := range tests {
		a, err := parseArrayHasAllValues(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*ArrayHasAllValuesX)(nil), a)
			assert.Equal(t, tt.left, a.Param)
			assert.Equal(t, tt.right, a.Slice)
			assert.Equal(t, tt.left, a.GetParam())
			assert.Equal(t, tt.right, a.GetValue())
			assert.Equal(t, tt.result, a.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestArrayHasValue_Equals(t *testing.T) {
	type testEqualsArrayHasValue struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testEqualsArrayHasValue{
		{
			e1:     ArrayHasValue(Param("a"), String("foo")),
			e2:     ArrayHasValue(Param("a"), String("foo")),
			result: true,
		},
		{
			e1:     ArrayHasValue(Param("a"), Integer(100)),
			e2:     ArrayHasValue(Param("a"), Integer(100)),
			result: true,
		},
		{
			e1:     ArrayHasValue(Param("a"), String("foo")),
			e2:     ArrayHasValue(Param("a"), String("bar")),
			result: false,
		},
		{
			e1:     ArrayHasValue(Param("a"), Integer(100)),
			e2:     ArrayHasValue(Param("a"), Integer(50)),
			result: false,
		},
		{
			e1:     ArrayHasValue(Param("a"), Integer(100)),
			e2:     ArrayHasValue(Param("a"), Float(100)),
			result: false,
		},
		{
			e1:     ArrayHasValue(Param("a"), Integer(100)),
			e2:     Integer(100),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}

func TestArrayHasAllValues_Equals(t *testing.T) {
	type testEqualsArrayHasAllValues struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testEqualsArrayHasAllValues{
		{
			e1:     ArrayHasAllValues(Param("a"), Slice(String("b"), String("c"), String("d"))),
			e2:     ArrayHasAllValues(Param("a"), Slice(String("b"), String("c"), String("d"))),
			result: true,
		},
		{
			e1:     ArrayHasAllValues(Param("a"), Slice(Integer(1), Integer(2), Integer(3))),
			e2:     ArrayHasAllValues(Param("a"), Slice(Integer(1), Integer(2), Integer(3))),
			result: true,
		},
		{
			e1:     ArrayHasAllValues(Param("a"), Slice(String("b"), String("c"), String("d"))),
			e2:     ArrayHasAllValues(Param("a"), Slice(String("b"), String("c"), String("f"))),
			result: false,
		},
		{
			e1:     ArrayHasAllValues(Param("a"), Slice(Integer(1), Integer(2), Integer(3))),
			e2:     ArrayHasAllValues(Param("a"), Slice(Integer(1), Integer(2), Integer(5))),
			result: false,
		},
		{
			e1:     ArrayHasAllValues(Param("a"), Slice(Integer(1), Integer(2), Integer(3))),
			e2:     ArrayHasAllValues(Param("a"), Slice(Integer(1), Integer(2), Float(3.25))),
			result: false,
		},
		{
			e1:     ArrayHasAllValues(Param("a"), Slice(String("b"), String("c"), String("d"))),
			e2:     String("a"),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}
