package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStatement(t *testing.T) {
	var (
		p = Param("a")
		v = Integer(1)
	)

	type testParseStatement struct {
		left  interface{}
		right interface{}
		err   error
	}
	var tests = []testParseStatement{
		{
			left:  p,
			right: v,
		},
		{
			left:  p,
			right: p,
		},
		{
			left:  nil,
			right: v,
			err:   IncorrectType("parseStatement", (*ParamX)(nil), nil),
		},
		{
			left:  p,
			right: nil,
			err:   IncorrectType("parseStatement", (*Value)(nil), nil),
		},
		{
			left:  v,
			right: v,
			err:   IncorrectType("parseStatement", (*ParamX)(nil), v),
		},
		{
			left:  (*Expression)(nil),
			right: v,
			err:   IncorrectType("parseStatement", (*ParamX)(nil), (*Expression)(nil)),
		},
		{
			left:  p,
			right: (*Expression)(nil),
			err:   IncorrectType("parseStatement", (*Value)(nil), (*Expression)(nil)),
		},
	}

	for _, tt := range tests {
		param, value, err := parseStatement(tt.left, tt.right)
		if tt.err == nil && assert.NoError(t, err) {
			assert.Equal(t, tt.left, param)
			assert.Equal(t, tt.right, value)
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}
