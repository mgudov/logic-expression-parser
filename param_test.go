package lep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseParam(t *testing.T) {
	type testParseParam struct {
		raw    []byte
		result string
	}
	var tests = []testParseParam{
		{
			raw:    []byte("a"),
			result: "a",
		},
		{
			raw:    []byte("    param_a    "),
			result: "param_a",
		},
		{
			raw:    []byte("object_a.field_a"),
			result: "object_a.field_a",
		},
	}

	for _, tt := range tests {
		p, err := parseParam(tt.raw)
		if assert.NoError(t, err) {
			assert.IsType(t, (*Param)(nil), p)
			assert.Equal(t, tt.result, p.Name)
			assert.Equal(t, tt.result, p.Value())
			assert.Equal(t, tt.result, p.String())
		}
	}
}

func TestParam_Equals(t *testing.T) {
	type testParamEquals struct {
		p1     Expression
		p2     Expression
		result bool
	}
	var tests = []testParamEquals{
		{
			p1:     NewParam("a"),
			p2:     NewParam("a"),
			result: true,
		},
		{
			p1:     NewParam("a"),
			p2:     NewParam("b"),
			result: false,
		},
		{
			p1:     NewParam("a"),
			p2:     NewString("a"),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.p1.Equals(tt.p2))
		assert.Equal(t, tt.result, tt.p2.Equals(tt.p1))
	}
}
