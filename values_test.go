package lep

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestParseString(t *testing.T) {
	type testParseString struct {
		raw    []byte
		result string
	}
	var tests = []testParseString{
		{
			raw:    []byte(`foo bar`),
			result: `foo bar`,
		},
		{
			raw:    []byte(`"foo bar"`),
			result: `foo bar`,
		},
		{
			raw:    []byte(`    "foo bar"    `),
			result: `foo bar`,
		},
		{
			raw:    []byte(`"!@#$%^&*()"`),
			result: `!@#$%^&*()`,
		},
	}

	for _, tt := range tests {
		v, err := parseString(tt.raw)
		if assert.NoError(t, err) {
			assert.IsType(t, (*StringX)(nil), v)
			assert.Equal(t, tt.result, v.Val)
			assert.Equal(t, tt.result, v.Value())
			assert.Equal(t, `"`+tt.result+`"`, v.String())
		}
	}
}

func TestParseInteger(t *testing.T) {
	type testParseInteger struct {
		raw    []byte
		result int64
		err    error
	}
	var tests = []testParseInteger{
		{
			raw:    []byte("1000"),
			result: 1000,
		},
		{
			raw:    []byte("    1000    "),
			result: 1000,
		},
		{
			raw:    []byte("-1000"),
			result: -1000,
		},
		{
			raw:    []byte("    -1000    "),
			result: -1000,
		},
		{
			raw:    []byte("not_integer"),
			result: 0,
			err: &strconv.NumError{
				Func: "ParseInt",
				Num:  "not_integer",
				Err:  strconv.ErrSyntax,
			},
		},
	}

	for _, tt := range tests {
		v, err := parseInteger(tt.raw)
		if tt.err == nil && assert.NoError(t, err) && v != nil {
			assert.IsType(t, (*IntegerX)(nil), v)
			assert.Equal(t, tt.result, v.Val)
			assert.Equal(t, tt.result, v.Value())
			assert.Equal(t, fmt.Sprintf("%d", tt.result), v.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseFloat(t *testing.T) {
	type testParseFloat struct {
		raw    []byte
		result float64
		err    error
	}
	var tests = []testParseFloat{
		{
			raw:    []byte("123.45"),
			result: 123.45,
		},
		{
			raw:    []byte("    123.45    "),
			result: 123.45,
		},
		{
			raw:    []byte("-123.45"),
			result: -123.45,
		},
		{
			raw:    []byte("    -123.45    "),
			result: -123.45,
		},
		{
			raw:    []byte("not_float"),
			result: 0,
			err: &strconv.NumError{
				Func: "ParseFloat",
				Num:  "not_float",
				Err:  strconv.ErrSyntax,
			},
		},
	}

	for _, tt := range tests {
		v, err := parseFloat(tt.raw)
		if tt.err == nil && assert.NoError(t, err) && v != nil {
			assert.IsType(t, (*FloatX)(nil), v)
			assert.Equal(t, tt.result, v.Val)
			assert.Equal(t, tt.result, v.Value())
			assert.Equal(t, strings.TrimRight(fmt.Sprintf("%f", tt.result), "0"), v.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseBoolean(t *testing.T) {
	type testParseBoolean struct {
		raw    []byte
		result bool
		err    error
	}
	var tests = []testParseBoolean{
		{
			raw:    []byte("true"),
			result: true,
		},
		{
			raw:    []byte("    true    "),
			result: true,
		},
		{
			raw:    []byte("false"),
			result: false,
		},
		{
			raw:    []byte("    false    "),
			result: false,
		},
		{
			raw: []byte("not_boolean"),
			err: IncorrectValue("parseBoolean", "(true/false)", "not_boolean"),
		},
	}

	for _, tt := range tests {
		v, err := parseBoolean(tt.raw)
		if tt.err == nil && assert.NoError(t, err) && v != nil {
			assert.IsType(t, (*BooleanX)(nil), v)
			assert.Equal(t, tt.result, v.Val)
			assert.Equal(t, tt.result, v.Value())
			assert.Equal(t, strconv.FormatBool(tt.result), v.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}

func TestParseNull(t *testing.T) {
	v, err := parseNull()
	if assert.NoError(t, err) {
		assert.IsType(t, (*NullX)(nil), v)
		assert.Equal(t, nil, v.Value())
		assert.Equal(t, "null", v.String())
	}
}

func TestString_Equals(t *testing.T) {
	type testStringEquals struct {
		v1     Expression
		v2     Expression
		result bool
	}
	var tests = []testStringEquals{
		{
			v1:     String("foo"),
			v2:     String("foo"),
			result: true,
		},
		{
			v1:     String("foo"),
			v2:     String("bar"),
			result: false,
		},
		{
			v1:     String("foo"),
			v2:     Integer(1),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.v1.Equals(tt.v2))
		assert.Equal(t, tt.result, tt.v2.Equals(tt.v1))
	}
}

func TestInteger_Equals(t *testing.T) {
	type testIntegerEquals struct {
		v1     Expression
		v2     Expression
		result bool
	}
	var tests = []testIntegerEquals{
		{
			v1:     Integer(100),
			v2:     Integer(100),
			result: true,
		},
		{
			v1:     Integer(-100),
			v2:     Integer(-100),
			result: true,
		},
		{
			v1:     Integer(100),
			v2:     Integer(200),
			result: false,
		},
		{
			v1:     Integer(-100),
			v2:     Integer(-200),
			result: false,
		},
		{
			v1:     Integer(100),
			v2:     Integer(-100),
			result: false,
		},
		{
			v1:     Integer(100),
			v2:     String("foo"),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.v1.Equals(tt.v2))
		assert.Equal(t, tt.result, tt.v2.Equals(tt.v1))
	}
}

func TestFloat_Equals(t *testing.T) {
	type testFloatEquals struct {
		v1     Expression
		v2     Expression
		result bool
	}
	var tests = []testFloatEquals{
		{
			v1:     Float(12.34),
			v2:     Float(12.34),
			result: true,
		},
		{
			v1:     Float(-12.34),
			v2:     Float(-12.34),
			result: true,
		},
		{
			v1:     Float(12.34),
			v2:     Float(12.35),
			result: false,
		},
		{
			v1:     Float(-12.34),
			v2:     Float(-12.35),
			result: false,
		},
		{
			v1:     Float(12.34),
			v2:     Float(-12.34),
			result: false,
		},
		{
			v1:     Float(12.34),
			v2:     String("foo"),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.v1.Equals(tt.v2))
		assert.Equal(t, tt.result, tt.v2.Equals(tt.v1))
	}
}

func TestBoolean_Equals(t *testing.T) {
	type testBooleanEquals struct {
		v1     Expression
		v2     Expression
		result bool
	}
	var tests = []testBooleanEquals{
		{
			v1:     Boolean(true),
			v2:     Boolean(true),
			result: true,
		},
		{
			v1:     Boolean(false),
			v2:     Boolean(false),
			result: true,
		},
		{
			v1:     Boolean(true),
			v2:     Boolean(false),
			result: false,
		},
		{
			v1:     Boolean(true),
			v2:     Integer(1),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.v1.Equals(tt.v2))
		assert.Equal(t, tt.result, tt.v2.Equals(tt.v1))
	}
}

func TestNull_Equals(t *testing.T) {
	type testNullEquals struct {
		v1     Expression
		v2     Expression
		result bool
	}
	var tests = []testNullEquals{
		{
			v1:     Null(),
			v2:     Null(),
			result: true,
		},
		{
			v1:     Null(),
			v2:     Boolean(false),
			result: false,
		},
		{
			v1:     Null(),
			v2:     Integer(0),
			result: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, tt.v1.Equals(tt.v2))
		assert.Equal(t, tt.result, tt.v2.Equals(tt.v1))
	}
}
