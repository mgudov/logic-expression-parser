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
