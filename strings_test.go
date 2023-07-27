package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStartsWith(t *testing.T) {
	type testParseStartsWith struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseStartsWith{
		{
			left:   Param("a"),
			right:  String("foo"),
			result: `a starts_with "foo"`,
		},
		{
			left:   Param("a"),
			right:  Param("b"),
			result: `a starts_with b`,
		},
		{
			left:  Param("a"),
			right: Float(0),
			err:   IncorrectType("parseStartsWith", (Stringify)(nil), (*FloatX)(nil)),
		},
		{
			left:  Integer(1),
			right: Integer(2),
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseStartsWith(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*StartsWithX)(nil), e)
			assert.Equal(t, tt.left, e.GetParam())
			assert.Equal(t, tt.right, e.GetValue())
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseEndsWith(t *testing.T) {
	type testParseEndsWith struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseEndsWith{
		{
			left:   Param("a"),
			right:  String("foo"),
			result: `a ends_with "foo"`,
		},
		{
			left:   Param("a"),
			right:  Param("b"),
			result: `a ends_with b`,
		},
		{
			left:  Param("a"),
			right: Float(0),
			err:   IncorrectType("parseEndsWith", (Stringify)(nil), (*FloatX)(nil)),
		},
		{
			left:  Integer(1),
			right: Integer(2),
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		e, err := parseEndsWith(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*EndsWithX)(nil), e)
			assert.Equal(t, tt.left, e.GetParam())
			assert.Equal(t, tt.right, e.GetValue())
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestStartsWith_Equals(t *testing.T) {
	type testStartsWithEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testStartsWithEquals{
		{
			e1:     StartsWith(Param("a"), String("foo")),
			e2:     StartsWith(Param("a"), String("foo")),
			result: true,
		},
		{
			e1:     StartsWith(Param("a"), String("foo")),
			e2:     StartsWith(Param("a"), Param("foo")),
			result: false,
		},
		{
			e1:     StartsWith(Param("a"), String("foo")),
			e2:     StartsWith(Param("a"), String("bar")),
			result: false,
		},
		{
			e1:     StartsWith(Param("a"), String("foo")),
			e2:     String("foo"),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}

func TestEndsWith_Equals(t *testing.T) {
	type testEndsWithEquals struct {
		e1     Expression
		e2     Expression
		result bool
	}
	var tests = []testEndsWithEquals{
		{
			e1:     EndsWith(Param("a"), String("foo")),
			e2:     EndsWith(Param("a"), String("foo")),
			result: true,
		},
		{
			e1:     EndsWith(Param("a"), String("foo")),
			e2:     EndsWith(Param("a"), Param("foo")),
			result: false,
		},
		{
			e1:     EndsWith(Param("a"), String("foo")),
			e2:     EndsWith(Param("a"), String("bar")),
			result: false,
		},
		{
			e1:     EndsWith(Param("a"), String("foo")),
			e2:     String("foo"),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.e1.Equals(tt.e2))
		assert.Equal(t, tt.result, tt.e2.Equals(tt.e1))
	}
}
