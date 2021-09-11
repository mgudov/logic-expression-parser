package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStatement(t *testing.T) {
	var (
		p = NewParam("a")
		v = NewInteger(1)
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
			err:   NewIncorrectType("parseStatement", (*Param)(nil), nil),
		},
		{
			left:  p,
			right: nil,
			err:   NewIncorrectType("parseStatement", (*Value)(nil), nil),
		},
		{
			left:  v,
			right: v,
			err:   NewIncorrectType("parseStatement", (*Param)(nil), v),
		},
		{
			left:  (*Expression)(nil),
			right: v,
			err:   NewIncorrectType("parseStatement", (*Param)(nil), (*Expression)(nil)),
		},
		{
			left:  p,
			right: (*Expression)(nil),
			err:   NewIncorrectType("parseStatement", (*Value)(nil), (*Expression)(nil)),
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
