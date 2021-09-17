package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSlice(t *testing.T) {
	var (
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testParseSlice struct {
		items  []interface{}
		result Value
	}
	var tests = []testParseSlice{
		{
			items:  []interface{}{v1, v2, v3, v4, v5},
			result: Slice(v1, v2, v3, v4, v5),
		},
		{
			items:  []interface{}{v1, "foo", v2, v3, "bar", v4, v5},
			result: Slice(v1, v2, v3, v4, v5),
		},
		{
			items:  []interface{}{v1, v3, []interface{}{v5}},
			result: Slice(v1, v3, v5),
		},
	}
	for _, tt := range tests {
		s, err := parseSlice(tt.items...)
		if assert.NoError(t, err) {
			assert.IsType(t, (*SliceX)(nil), s)
			assert.Equal(t, tt.result, s)
			assert.Equal(t, tt.result.Value(), s.Value())
			assert.Equal(t, tt.result.String(), s.String())
		}
	}
}

func TestParseInSlice(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testParseInSlice struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseInSlice{
		{
			left:   p1,
			right:  Slice(v1, v2, v3, v4, v5),
			result: `a in ["foo",100,12.34,true,null]`,
		},
		{
			left:  p1,
			right: v1,
			err:   IncorrectType("parseInSlice", (*SliceX)(nil), (*StringX)(nil)),
		},
		{
			left:  Param("a"),
			right: Equals(p1, v1),
			err:   IncorrectType("parseStatement", (*Value)(nil), (*EqualsX)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseInSlice(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*InSliceX)(nil), s)
			assert.Equal(t, tt.left, s.Param)
			assert.Equal(t, tt.right, s.Slice)
			assert.Equal(t, tt.result, s.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseNotInSlice(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testParseNotInSlice struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseNotInSlice{
		{
			left:   p1,
			right:  Slice(v1, v2, v3, v4, v5),
			result: `a not_in ["foo",100,12.34,true,null]`,
		},
		{
			left:  p1,
			right: v1,
			err:   IncorrectType("parseNotInSlice", (*SliceX)(nil), (*StringX)(nil)),
		},
		{
			left:  Param("a"),
			right: Equals(p1, v1),
			err:   IncorrectType("parseStatement", (*Value)(nil), (*EqualsX)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseNotInSlice(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*NotInSliceX)(nil), s)
			assert.Equal(t, tt.left, s.Param)
			assert.Equal(t, tt.right, s.Slice)
			assert.Equal(t, tt.result, s.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestSlice_Equals(t *testing.T) {
	var (
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testSliceEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testSliceEquals{
		{
			s1:     Slice(v1, v2, v3, v4, v5),
			s2:     Slice(v1, v2, v3, v4, v5),
			result: true,
		},
		{
			s1:     Slice(v1, v2, v3, v4, v5),
			s2:     Slice(v5, v4, v3, v2, v1),
			result: false,
		},
		{
			s1:     Slice(v1, v2, v3, v4, v5),
			s2:     Slice(v1, v2, v3),
			result: false,
		},
		{
			s1:     Slice(v1, v2, v3, v4, v5),
			s2:     v1,
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}

func TestInSlice_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testInSliceEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testInSliceEquals{
		{
			s1:     InSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     InSlice(p1, Slice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     InSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     InSlice(p1, Slice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     InSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     InSlice(p1, Slice(v1, v2, v3)),
			result: false,
		},
		{
			s1:     InSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Slice(v1, v2, v3, v4, v5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}

func TestNotInSlice_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testNotInSliceEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testNotInSliceEquals{
		{
			s1:     NotInSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     NotInSlice(p1, Slice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     NotInSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     NotInSlice(p1, Slice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     NotInSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     NotInSlice(p1, Slice(v1, v2, v3)),
			result: false,
		},
		{
			s1:     NotInSlice(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Slice(v1, v2, v3, v4, v5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}
