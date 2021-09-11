package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSlice(t *testing.T) {
	var (
		v1 = NewString("foo")
		v2 = NewInteger(100)
		v3 = NewFloat(12.34)
		v4 = NewBoolean(true)
		v5 = NewNull()
	)

	type testParseSlice struct {
		items  []interface{}
		result Value
	}
	var tests = []testParseSlice{
		{
			items:  []interface{}{v1, v2, v3, v4, v5},
			result: NewSlice(v1, v2, v3, v4, v5),
		},
		{
			items:  []interface{}{v1, "foo", v2, v3, "bar", v4, v5},
			result: NewSlice(v1, v2, v3, v4, v5),
		},
		{
			items:  []interface{}{v1, v3, []interface{}{v5}},
			result: NewSlice(v1, v3, v5),
		},
	}
	for _, tt := range tests {
		s, err := parseSlice(tt.items...)
		if assert.NoError(t, err) {
			assert.IsType(t, (*Slice)(nil), s)
			assert.Equal(t, tt.result, s)
			assert.Equal(t, tt.result.Value(), s.Value())
			assert.Equal(t, tt.result.String(), s.String())
		}
	}
}

func TestParseInSlice(t *testing.T) {
	var (
		p1 = NewParam("a")
		v1 = NewString("foo")
		v2 = NewInteger(100)
		v3 = NewFloat(12.34)
		v4 = NewBoolean(true)
		v5 = NewNull()
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
			right:  NewSlice(v1, v2, v3, v4, v5),
			result: `a in ["foo",100,12.34,true,null]`,
		},
		{
			left:  p1,
			right: v1,
			err:   NewIncorrectType("parseInSlice", (*Slice)(nil), (*String)(nil)),
		},
		{
			left:  NewParam("a"),
			right: NewEquals(p1, v1),
			err:   NewIncorrectType("parseStatement", (*Value)(nil), (*Equals)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseInSlice(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && s != nil {
			assert.IsType(t, (*InSlice)(nil), s)
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
		p1 = NewParam("a")
		v1 = NewString("foo")
		v2 = NewInteger(100)
		v3 = NewFloat(12.34)
		v4 = NewBoolean(true)
		v5 = NewNull()
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
			right:  NewSlice(v1, v2, v3, v4, v5),
			result: `a not_in ["foo",100,12.34,true,null]`,
		},
		{
			left:  p1,
			right: v1,
			err:   NewIncorrectType("parseNotInSlice", (*Slice)(nil), (*String)(nil)),
		},
		{
			left:  NewParam("a"),
			right: NewEquals(p1, v1),
			err:   NewIncorrectType("parseStatement", (*Value)(nil), (*Equals)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseNotInSlice(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) && s != nil {
			assert.IsType(t, (*NotInSlice)(nil), s)
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
		v1 = NewString("foo")
		v2 = NewInteger(100)
		v3 = NewFloat(12.34)
		v4 = NewBoolean(true)
		v5 = NewNull()
	)

	type testSliceEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testSliceEquals{
		{
			s1:     NewSlice(v1, v2, v3, v4, v5),
			s2:     NewSlice(v1, v2, v3, v4, v5),
			result: true,
		},
		{
			s1:     NewSlice(v1, v2, v3, v4, v5),
			s2:     NewSlice(v5, v4, v3, v2, v1),
			result: false,
		},
		{
			s1:     NewSlice(v1, v2, v3, v4, v5),
			s2:     NewSlice(v1, v2, v3),
			result: false,
		},
		{
			s1:     NewSlice(v1, v2, v3, v4, v5),
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
		p1 = NewParam("a")
		v1 = NewString("foo")
		v2 = NewInteger(100)
		v3 = NewFloat(12.34)
		v4 = NewBoolean(true)
		v5 = NewNull()
	)

	type testInSliceEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testInSliceEquals{
		{
			s1:     NewInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     NewInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewInSlice(p1, NewSlice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     NewInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewInSlice(p1, NewSlice(v1, v2, v3)),
			result: false,
		},
		{
			s1:     NewInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewSlice(v1, v2, v3, v4, v5),
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
		p1 = NewParam("a")
		v1 = NewString("foo")
		v2 = NewInteger(100)
		v3 = NewFloat(12.34)
		v4 = NewBoolean(true)
		v5 = NewNull()
	)

	type testNotInSliceEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testNotInSliceEquals{
		{
			s1:     NewNotInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewNotInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     NewNotInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewNotInSlice(p1, NewSlice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     NewNotInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewNotInSlice(p1, NewSlice(v1, v2, v3)),
			result: false,
		},
		{
			s1:     NewNotInSlice(p1, NewSlice(v1, v2, v3, v4, v5)),
			s2:     NewSlice(v1, v2, v3, v4, v5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}
