package lep

import "strings"

type Param struct {
	Name string
}

var _ Value = (*Param)(nil)

func NewParam(name string) *Param {
	return &Param{Name: name}
}

func (p Param) Equals(other Expression) bool {
	if expr, ok := other.(*Param); ok {
		return p.Name == expr.Name
	}
	return false
}

func (p Param) String() string {
	return p.Name
}

func (p Param) Value() interface{} {
	return p.Name
}

func parseParam(b []byte) (*Param, error) {
	return NewParam(strings.TrimSpace(string(b))), nil
}
