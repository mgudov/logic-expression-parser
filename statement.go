package lep

type Statement interface {
	GetParam() *ParamX
	GetValue() Value
}

type statement struct {
	Param *ParamX
	Value Value
}

var _ Statement = (*statement)(nil)

func (s statement) GetParam() *ParamX {
	return s.Param
}

func (s statement) GetValue() Value {
	return s.Value
}

func parseStatement(left, right interface{}) (*ParamX, Value, error) {
	param, ok := left.(*ParamX)
	if !ok {
		return nil, nil, IncorrectType("parseStatement", (*ParamX)(nil), left)
	}
	value, ok := right.(Value)
	if !ok {
		return nil, nil, IncorrectType("parseStatement", (*Value)(nil), right)
	}
	return param, value, nil
}
