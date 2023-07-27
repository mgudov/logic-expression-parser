package lep

import (
	"fmt"
	"strconv"
	"strings"
)

type StringX struct {
	Val string
}

var _ Stringify = (*StringX)(nil)

func String(val string) *StringX {
	return &StringX{Val: val}
}

func (s StringX) Equals(other Expression) bool {
	if expr, ok := other.(*StringX); ok {
		return s.Val == expr.Val
	}
	return false
}

func (s StringX) String() string {
	return `"` + s.Val + `"`
}

func (s StringX) Value() interface{} {
	return s.Val
}

func (s StringX) IsStringify() bool {
	return true
}

func parseString(b []byte) (*StringX, error) {
	val := strings.TrimSpace(string(b))
	val = strings.TrimPrefix(val, `"`)
	val = strings.TrimSuffix(val, `"`)
	return String(val), nil
}

type IntegerX struct {
	Val int64
}

var _ Value = (*IntegerX)(nil)

func Integer(val int64) *IntegerX {
	return &IntegerX{Val: val}
}

func (i IntegerX) Equals(other Expression) bool {
	if expr, ok := other.(*IntegerX); ok {
		return i.Val == expr.Val
	}
	return false
}

func (i IntegerX) String() string {
	return fmt.Sprintf("%d", i.Val)
}

func (i IntegerX) Value() interface{} {
	return i.Val
}

func parseInteger(b []byte) (*IntegerX, error) {
	val, err := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return nil, err
	}
	return Integer(val), nil
}

type FloatX struct {
	Val float64
}

var _ Value = (*FloatX)(nil)

func Float(val float64) *FloatX {
	return &FloatX{Val: val}
}

func (f FloatX) Equals(other Expression) bool {
	if expr, ok := other.(*FloatX); ok {
		return f.Val == expr.Val
	}
	return false
}

func (f FloatX) String() string {
	return strings.TrimRight(fmt.Sprintf("%f", f.Val), "0")
}

func (f FloatX) Value() interface{} {
	return f.Val
}

func parseFloat(b []byte) (*FloatX, error) {
	val, err := strconv.ParseFloat(strings.TrimSpace(string(b)), 64)
	if err != nil {
		return nil, err
	}
	return Float(val), nil
}

type BooleanX struct {
	Val bool
}

var _ Value = (*BooleanX)(nil)

func Boolean(val bool) *BooleanX {
	return &BooleanX{Val: val}
}

func (v BooleanX) Equals(other Expression) bool {
	if expr, ok := other.(*BooleanX); ok {
		return v.Val == expr.Val
	}
	return false
}

func (v BooleanX) String() string {
	if v.Val {
		return "true"
	} else {
		return "false"
	}
}

func parseBoolean(b []byte) (*BooleanX, error) {
	switch val := strings.TrimSpace(string(b)); val {
	default:
		return nil, IncorrectValue("parseBoolean", "(true/false)", val)
	case "true":
		return Boolean(true), nil
	case "false":
		return Boolean(false), nil
	}
}

func (v BooleanX) Value() interface{} {
	return v.Val
}

type NullX struct{}

var _ Value = (*NullX)(nil)

func Null() *NullX {
	return &NullX{}
}

func (NullX) Equals(other Expression) bool {
	_, ok := other.(*NullX)
	return ok
}

func (NullX) String() string {
	return "null"
}

func (NullX) Value() interface{} {
	return nil
}

func parseNull() (*NullX, error) {
	return Null(), nil
}
