package lep

import (
	"strings"
)

type Slice struct {
	Values []Value
}

var _ Value = (*Slice)(nil)

func NewSlice(values ...Value) *Slice {
	return &Slice{Values: values}
}

func (e Slice) Equals(other Expression) bool {
	if expr, ok := other.(*Slice); ok {
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

func (e Slice) String() string {
	var items []string
	for _, value := range e.Values {
		items = append(items, value.String())
	}
	return "[" + strings.Join(items, ",") + "]"
}

func (e Slice) Value() interface{} {
	return e.Values
}

func parseSlice(items ...interface{}) (*Slice, error) {
	var values []Value
	for _, e := range parseExpressions(items...) {
		if v, ok := e.(Value); ok {
			values = append(values, v)
		}
	}
	return NewSlice(values...), nil
}

type InSlice struct {
	Param *Param
	Slice *Slice
}

var _ Expression = (*InSlice)(nil)

func NewInSlice(param *Param, slice *Slice) *InSlice {
	return &InSlice{
		Param: param,
		Slice: slice,
	}
}

func (e InSlice) Equals(other Expression) bool {
	if expr, ok := other.(*InSlice); ok {
		return e.Param.Equals(expr.Param) && e.Slice.Equals(expr.Slice)
	}
	return false
}

func (e InSlice) String() string {
	return e.Param.String() + " in " + e.Slice.String()
}

func parseInSlice(left, right interface{}) (*InSlice, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	slice, ok := value.(*Slice)
	if !ok {
		return nil, NewIncorrectType("parseInSlice", (*Slice)(nil), value)
	}
	return NewInSlice(param, slice), nil
}

type NotInSlice struct {
	Param *Param
	Slice *Slice
}

var _ Expression = (*NotInSlice)(nil)

func NewNotInSlice(param *Param, slice *Slice) *NotInSlice {
	return &NotInSlice{
		Param: param,
		Slice: slice,
	}
}

func (e NotInSlice) Equals(other Expression) bool {
	if expr, ok := other.(*NotInSlice); ok {
		return e.Param.Equals(expr.Param) && e.Slice.Equals(expr.Slice)
	}
	return false
}

func (e NotInSlice) String() string {
	return e.Param.String() + " not_in " + e.Slice.String()
}

func parseNotInSlice(left, right interface{}) (*NotInSlice, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	slice, ok := value.(*Slice)
	if !ok {
		return nil, NewIncorrectType("parseNotInSlice", (*Slice)(nil), value)
	}
	return NewNotInSlice(param, slice), nil
}
