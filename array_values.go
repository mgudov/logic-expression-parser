package lep

import "fmt"

type ArrayHasValueX struct {
	statement
}

var _ Expression = (*ArrayHasValueX)(nil)
var _ Statement = (*ArrayHasValueX)(nil)

func ArrayHasValue(param *ParamX, value Value) *ArrayHasValueX {
	return &ArrayHasValueX{
		statement: statement{
			Param: param,
			Value: value,
		},
	}
}

func (a ArrayHasValueX) Equals(other Expression) bool {
	if expr, ok := other.(*ArrayHasValueX); ok {
		return a.Param.Equals(expr.Param) && a.Value.Equals(expr.Value)
	}
	return false
}

func (a ArrayHasValueX) String() string {
	return fmt.Sprintf("%s %s %s", a.Param.String(), "has", a.Value.String())
}

func parseArrayHasValue(left, right interface{}) (*ArrayHasValueX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return ArrayHasValue(param, value), nil
}

type ArrayHasAllValuesX struct {
	Param *ParamX
	Slice *SliceX
}

func ArrayHasAllValues(param *ParamX, slice *SliceX) *ArrayHasAllValuesX {
	return &ArrayHasAllValuesX{
		Param: param,
		Slice: slice,
	}
}

func (a ArrayHasAllValuesX) Equals(other Expression) bool {
	if expr, ok := other.(*ArrayHasAllValuesX); ok {
		return a.Param.Equals(expr.Param) && a.Slice.Equals(expr.Slice)
	}
	return false
}

func (a ArrayHasAllValuesX) String() string {
	return a.Param.String() + " has_all " + a.Slice.String()
}

func (a ArrayHasAllValuesX) GetParam() *ParamX {
	return a.Param
}

func (a ArrayHasAllValuesX) GetValue() Value {
	return a.Slice
}

func parseArrayHasAllValues(left, right interface{}) (*ArrayHasAllValuesX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	slice, ok := value.(*SliceX)
	if !ok {
		return nil, IncorrectType("parseArrayHasAllValues", (*SliceX)(nil), value)
	}
	return ArrayHasAllValues(param, slice), nil
}
