package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHas(t *testing.T) {
	var (
		p1 = Param("a")
		p2 = Param("b")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testParseHas struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseHas{
		{
			left:   p1,
			right:  v1,
			result: `a has "foo"`,
		},
		{
			left:   p1,
			right:  v2,
			result: `a has 100`,
		},
		{
			left:   p1,
			right:  v3,
			result: `a has 12.34`,
		},
		{
			left:   p1,
			right:  p2,
			result: `a has b`,
		},
		{
			left:   p1,
			right:  Slice(v1, v2, v3, v4, v5),
			result: `a has ["foo",100,12.34,true,null]`,
		},
		{
			left:  v1,
			right: p1,
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*StringX)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseHas(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*HasX)(nil), s)
			assert.Equal(t, tt.left, s.Param)
			assert.Equal(t, tt.right, s.Value)
			assert.Equal(t, tt.left, s.GetParam())
			assert.Equal(t, tt.right, s.GetValue())
			assert.Equal(t, tt.result, s.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseNotHas(t *testing.T) {
	var (
		p1 = Param("a")
		p2 = Param("b")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testParseNotHas struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseNotHas{
		{
			left:   p1,
			right:  v1,
			result: `a not_has "foo"`,
		},
		{
			left:   p1,
			right:  v2,
			result: `a not_has 100`,
		},
		{
			left:   p1,
			right:  v3,
			result: `a not_has 12.34`,
		},
		{
			left:   p1,
			right:  p2,
			result: `a not_has b`,
		},
		{
			left:   p1,
			right:  Slice(v1, v2, v3, v4, v5),
			result: `a not_has ["foo",100,12.34,true,null]`,
		},
		{
			left:  v1,
			right: p1,
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*StringX)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseNotHas(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*NotHasX)(nil), s)
			assert.Equal(t, tt.left, s.Param)
			assert.Equal(t, tt.right, s.Value)
			assert.Equal(t, tt.left, s.GetParam())
			assert.Equal(t, tt.right, s.GetValue())
			assert.Equal(t, tt.result, s.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseHasAny(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testParseHasAny struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseHasAny{
		{
			left:   p1,
			right:  Slice(v1, v2, v3, v4, v5),
			result: `a has_any ["foo",100,12.34,true,null]`,
		},
		{
			left:  p1,
			right: v1,
			err:   IncorrectType("parseHasAny", (*SliceX)(nil), (*StringX)(nil)),
		},
		{
			left:  Param("a"),
			right: Equals(p1, v1),
			err:   IncorrectType("parseStatement", (*Value)(nil), (*EqualsX)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseHasAny(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*HasAnyX)(nil), s)
			assert.Equal(t, tt.left, s.Param)
			assert.Equal(t, tt.right, s.Slice)
			assert.Equal(t, tt.left, s.GetParam())
			assert.Equal(t, tt.right, s.GetValue())
			assert.Equal(t, tt.result, s.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseHasAll(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testParseHasAll struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseHasAll{
		{
			left:   p1,
			right:  Slice(v1, v2, v3, v4, v5),
			result: `a has_all ["foo",100,12.34,true,null]`,
		},
		{
			left:  p1,
			right: v1,
			err:   IncorrectType("parseHasAll", (*SliceX)(nil), (*StringX)(nil)),
		},
		{
			left:  Param("a"),
			right: Equals(p1, v1),
			err:   IncorrectType("parseStatement", (*Value)(nil), (*EqualsX)(nil)),
		},
	}

	for _, tt := range tests {
		s, err := parseHasAll(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*HasAllX)(nil), s)
			assert.Equal(t, tt.left, s.Param)
			assert.Equal(t, tt.right, s.Slice)
			assert.Equal(t, tt.left, s.GetParam())
			assert.Equal(t, tt.right, s.GetValue())
			assert.Equal(t, tt.result, s.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestHas_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testHasEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testHasEquals{
		{
			s1:     Has(p1, v1),
			s2:     Has(p1, v1),
			result: true,
		},
		{
			s1:     Has(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Has(p1, Slice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     Has(p1, v1),
			s2:     Has(p1, v2),
			result: false,
		},
		{
			s1:     Has(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Has(p1, Slice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     Has(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Slice(v1, v2, v3, v4, v5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}

func TestNotHas_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testHasEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testHasEquals{
		{
			s1:     NotHas(p1, v1),
			s2:     NotHas(p1, v1),
			result: true,
		},
		{
			s1:     NotHas(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     NotHas(p1, Slice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     NotHas(p1, v1),
			s2:     NotHas(p1, v2),
			result: false,
		},
		{
			s1:     NotHas(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     NotHas(p1, Slice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     NotHas(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Slice(v1, v2, v3, v4, v5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}

func TestHasAny_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testHasAnyEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testHasAnyEquals{
		{
			s1:     HasAny(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     HasAny(p1, Slice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     HasAny(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     HasAny(p1, Slice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     HasAny(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Slice(v1, v2, v3, v4, v5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}

func TestHasAll_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		v1 = String("foo")
		v2 = Integer(100)
		v3 = Float(12.34)
		v4 = Boolean(true)
		v5 = Null()
	)

	type testHasAllEquals struct {
		s1     Expression
		s2     Expression
		result bool
	}
	var tests = []testHasAllEquals{
		{
			s1:     HasAll(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     HasAll(p1, Slice(v1, v2, v3, v4, v5)),
			result: true,
		},
		{
			s1:     HasAll(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     HasAll(p1, Slice(v5, v4, v3, v2, v1)),
			result: false,
		},
		{
			s1:     HasAll(p1, Slice(v1, v2, v3, v4, v5)),
			s2:     Slice(v1, v2, v3, v4, v5),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.s1.Equals(tt.s2))
		assert.Equal(t, tt.result, tt.s2.Equals(tt.s1))
	}
}
