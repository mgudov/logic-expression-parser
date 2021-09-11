package lep

import (
	"fmt"
	"strconv"
	"strings"
)

type String struct {
	Val string
}

var _ Value = (*String)(nil)

func NewString(val string) *String {
	return &String{Val: val}
}

func (s String) Equals(other Expression) bool {
	if expr, ok := other.(*String); ok {
		return s.Val == expr.Val
	}
	return false
}

func (s String) String() string {
	return `"` + s.Val + `"`
}

func (s String) Value() interface{} {
	return s.Val
}

func parseString(b []byte) (*String, error) {
	return NewString(strings.Trim(string(b), ` "`)), nil
}

type Integer struct {
	Val int64
}

var _ Value = (*Integer)(nil)

func NewInteger(val int64) *Integer {
	return &Integer{Val: val}
}

func (i Integer) Equals(other Expression) bool {
	if expr, ok := other.(*Integer); ok {
		return i.Val == expr.Val
	}
	return false
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i.Val)
}

func (i Integer) Value() interface{} {
	return i.Val
}

func parseInteger(b []byte) (*Integer, error) {
	val, err := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return nil, err
	}
	return NewInteger(val), nil
}

type Float struct {
	Val float64
}

var _ Value = (*Float)(nil)

func NewFloat(val float64) *Float {
	return &Float{Val: val}
}

func (f Float) Equals(other Expression) bool {
	if expr, ok := other.(*Float); ok {
		return f.Val == expr.Val
	}
	return false
}

func (f Float) String() string {
	return strings.TrimRight(fmt.Sprintf("%f", f.Val), "0")
}

func (f Float) Value() interface{} {
	return f.Val
}

func parseFloat(b []byte) (*Float, error) {
	val, err := strconv.ParseFloat(strings.TrimSpace(string(b)), 64)
	if err != nil {
		return nil, err
	}
	return NewFloat(val), nil
}

type Boolean struct {
	Val bool
}

var _ Value = (*Boolean)(nil)

func NewBoolean(val bool) *Boolean {
	return &Boolean{Val: val}
}

func (v Boolean) Equals(other Expression) bool {
	if expr, ok := other.(*Boolean); ok {
		return v.Val == expr.Val
	}
	return false
}

func (v Boolean) String() string {
	if v.Val {
		return "true"
	} else {
		return "false"
	}
}

func parseBoolean(b []byte) (*Boolean, error) {
	switch val := strings.TrimSpace(string(b)); val {
	default:
		return nil, NewIncorrectValue("parseBoolean", "(true/false)", val)
	case "true":
		return NewBoolean(true), nil
	case "false":
		return NewBoolean(false), nil
	}
}

func (v Boolean) Value() interface{} {
	return v.Val
}

type Null struct{}

var _ Value = (*Null)(nil)

func NewNull() *Null {
	return &Null{}
}

func (Null) Equals(other Expression) bool {
	_, ok := other.(*Null)
	return ok
}

func (Null) String() string {
	return "null"
}

func (Null) Value() interface{} {
	return nil
}

func parseNull() (*Null, error) {
	return NewNull(), nil
}
