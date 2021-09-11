package lep

import "strings"

type ParamX struct {
	Name string
}

var _ Value = (*ParamX)(nil)

func Param(name string) *ParamX {
	return &ParamX{Name: name}
}

func (p ParamX) Equals(other Expression) bool {
	if expr, ok := other.(*ParamX); ok {
		return p.Name == expr.Name
	}
	return false
}

func (p ParamX) String() string {
	return p.Name
}

func (p ParamX) Value() interface{} {
	return p.Name
}

func parseParam(b []byte) (*ParamX, error) {
	return Param(strings.TrimSpace(string(b))), nil
}
