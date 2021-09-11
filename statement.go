package lep

type statement struct {
	Param *ParamX
	Value Value
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
