package lep

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"regexp/syntax"
	"testing"
)

func TestParseRegexp(t *testing.T) {
	type testParseRegexp struct {
		raw    []byte
		result *regexp.Regexp
		err    error
	}
	var tests = []testParseRegexp{
		{
			raw:    []byte("/[a-zA-Z0-9]+/"),
			result: regexp.MustCompile("/[a-zA-Z0-9]+/"),
		},
		{
			raw: []byte("/([a-z]+/"),
			err: &syntax.Error{Code: syntax.ErrMissingParen, Expr: "/([a-z]+/"},
		},
	}

	for _, tt := range tests {
		r, err := parseRegexp(tt.raw)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*RegexpX)(nil), r)
			assert.Equal(t, tt.result, r.Regexp)
			assert.Equal(t, tt.result.String(), r.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseMatchRegexp(t *testing.T) {
	type testParseMatchRegexp struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseMatchRegexp{
		{
			left:   Param("a"),
			right:  Regexp(regexp.MustCompile(`/[a-z]+/gm`)),
			result: "a =~ /[a-z]+/gm",
		},
		{
			left:  Integer(1),
			right: Regexp(regexp.MustCompile(`/[a-z]+/gm`)),
			err:   IncorrectType("parseMatchRegexp", (*ParamX)(nil), (*IntegerX)(nil)),
		},
		{
			left:  Param("a"),
			right: Integer(1),
			err:   IncorrectType("parseMatchRegexp", (*RegexpX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		r, err := parseMatchRegexp(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*MatchRegexpX)(nil), r)
			assert.Equal(t, tt.left, r.Param)
			assert.Equal(t, tt.right, r.Regexp)
			assert.Equal(t, tt.result, r.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseNotMatchRegexp(t *testing.T) {
	type testParseNotMatchRegexp struct {
		left   interface{}
		right  interface{}
		result string
		err    error
	}
	var tests = []testParseNotMatchRegexp{
		{
			left:   Param("a"),
			right:  Regexp(regexp.MustCompile(`/[a-z]+/gm`)),
			result: "a =~ /[a-z]+/gm",
		},
		{
			left:  Integer(1),
			right: Regexp(regexp.MustCompile(`/[a-z]+/gm`)),
			err:   IncorrectType("parseNotMatchRegexp", (*ParamX)(nil), (*IntegerX)(nil)),
		},
		{
			left:  Param("a"),
			right: Integer(1),
			err:   IncorrectType("parseNotMatchRegexp", (*RegexpX)(nil), (*IntegerX)(nil)),
		},
	}

	for _, tt := range tests {
		r, err := parseNotMatchRegexp(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*NotMatchRegexpX)(nil), r)
			assert.Equal(t, tt.left, r.Param)
			assert.Equal(t, tt.right, r.Regexp)
			assert.Equal(t, tt.result, r.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestRegexp_Equals(t *testing.T) {
	var (
		re1 = regexp.MustCompile("/[a-zA-Z0-9]+/")
		re2 = regexp.MustCompile("/[a-zA-Z]+/")
	)

	type testRegexpEquals struct {
		r1     Expression
		r2     Expression
		result bool
	}
	var tests = []testRegexpEquals{
		{
			r1:     Regexp(re1),
			r2:     Regexp(re1),
			result: true,
		},
		{
			r1:     Regexp(re1),
			r2:     Regexp(re2),
			result: false,
		},
		{
			r1:     Regexp(re1),
			r2:     Integer(1),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.r1.Equals(tt.r2))
		assert.Equal(t, tt.result, tt.r2.Equals(tt.r1))
	}
}

func TestMatchRegexp_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		p2 = Param("b")
		r1 = Regexp(regexp.MustCompile("/[a-zA-Z0-9]+/"))
		r2 = Regexp(regexp.MustCompile("/[a-zA-Z]+/"))
	)

	type testMatchRegexp struct {
		r1     Expression
		r2     Expression
		result bool
	}
	var tests = []testMatchRegexp{
		{
			r1:     MatchRegexp(p1, r1),
			r2:     MatchRegexp(p1, r1),
			result: true,
		},
		{
			r1:     MatchRegexp(p1, r1),
			r2:     MatchRegexp(p1, r2),
			result: false,
		},
		{
			r1:     MatchRegexp(p1, r1),
			r2:     MatchRegexp(p2, r1),
			result: false,
		},
		{
			r1:     MatchRegexp(p2, r1),
			r2:     MatchRegexp(p2, r2),
			result: false,
		},
		{
			r1:     MatchRegexp(p2, r1),
			r2:     r2,
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.r1.Equals(tt.r2))
		assert.Equal(t, tt.result, tt.r2.Equals(tt.r1))
	}
}

func TestNotMatchRegexp_Equals(t *testing.T) {
	var (
		p1 = Param("a")
		p2 = Param("b")
		r1 = Regexp(regexp.MustCompile("/[a-zA-Z0-9]+/"))
		r2 = Regexp(regexp.MustCompile("/[a-zA-Z]+/"))
	)

	type testNotMatchRegexp struct {
		r1     Expression
		r2     Expression
		result bool
	}
	var tests = []testNotMatchRegexp{
		{
			r1:     NotMatchRegexp(p1, r1),
			r2:     NotMatchRegexp(p1, r1),
			result: true,
		},
		{
			r1:     NotMatchRegexp(p1, r1),
			r2:     NotMatchRegexp(p1, r2),
			result: false,
		},
		{
			r1:     NotMatchRegexp(p1, r1),
			r2:     NotMatchRegexp(p2, r1),
			result: false,
		},
		{
			r1:     NotMatchRegexp(p2, r1),
			r2:     NotMatchRegexp(p2, r2),
			result: false,
		},
		{
			r1:     NotMatchRegexp(p2, r1),
			r2:     r2,
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.r1.Equals(tt.r2))
		assert.Equal(t, tt.result, tt.r2.Equals(tt.r1))
	}
}
