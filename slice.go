package lep

import (
	"strings"
)

type SliceX struct {
	Values []Value
}

var _ Value = (*SliceX)(nil)

func Slice(values ...Value) *SliceX {
	return &SliceX{Values: values}
}

func (e SliceX) Equals(other Expression) bool {
	if expr, ok := other.(*SliceX); ok {
		if len(e.Values) != len(expr.Values) {
			return false
		}
		for i, value := range e.Values {
			if !value.Equals(expr.Values[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (e SliceX) String() string {
	var items []string
	for _, value := range e.Values {
		items = append(items, value.String())
	}
	return "[" + strings.Join(items, ",") + "]"
}

func (e SliceX) Value() interface{} {
	return e.Values
}

func parseSlice(items ...interface{}) (*SliceX, error) {
	var values []Value
	for _, e := range parseExpressions(items...) {
		if v, ok := e.(Value); ok {
			values = append(values, v)
		}
	}
	return Slice(values...), nil
}

type InSliceX struct {
	Param *ParamX
	Slice *SliceX
}

var _ Expression = (*InSliceX)(nil)
var _ Statement = (*InSliceX)(nil)

func InSlice(param *ParamX, slice *SliceX) *InSliceX {
	return &InSliceX{
		Param: param,
		Slice: slice,
	}
}

func (e InSliceX) Equals(other Expression) bool {
	if expr, ok := other.(*InSliceX); ok {
		return e.Param.Equals(expr.Param) && e.Slice.Equals(expr.Slice)
	}
	return false
}

func (e InSliceX) String() string {
	return e.Param.String() + " in " + e.Slice.String()
}

func (e InSliceX) GetParam() *ParamX {
	return e.Param
}

func (e InSliceX) GetValue() Value {
	return e.Slice
}

func parseInSlice(left, right interface{}) (*InSliceX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	slice, ok := value.(*SliceX)
	if !ok {
		return nil, IncorrectType("parseInSlice", (*SliceX)(nil), value)
	}
	return InSlice(param, slice), nil
}

type NotInSliceX struct {
	Param *ParamX
	Slice *SliceX
}

var _ Expression = (*NotInSliceX)(nil)
var _ Statement = (*NotInSliceX)(nil)

func NotInSlice(param *ParamX, slice *SliceX) *NotInSliceX {
	return &NotInSliceX{
		Param: param,
		Slice: slice,
	}
}

func (e NotInSliceX) Equals(other Expression) bool {
	if expr, ok := other.(*NotInSliceX); ok {
		return e.Param.Equals(expr.Param) && e.Slice.Equals(expr.Slice)
	}
	return false
}

func (e NotInSliceX) String() string {
	return e.Param.String() + " not_in " + e.Slice.String()
}

func (e NotInSliceX) GetParam() *ParamX {
	return e.Param
}

func (e NotInSliceX) GetValue() Value {
	return e.Slice
}

func parseNotInSlice(left, right interface{}) (*NotInSliceX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	slice, ok := value.(*SliceX)
	if !ok {
		return nil, IncorrectType("parseNotInSlice", (*SliceX)(nil), value)
	}
	return NotInSlice(param, slice), nil
}
