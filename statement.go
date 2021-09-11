package lep

type statement struct {
	Param *Param
	Value Value
}

func parseStatement(left, right interface{}) (*Param, Value, error) {
	param, ok := left.(*Param)
	if !ok {
		return nil, nil, NewIncorrectType("parseStatement", (*Param)(nil), left)
	}
	value, ok := right.(Value)
	if !ok {
		return nil, nil, NewIncorrectType("parseStatement", (*Value)(nil), right)
	}
	return param, value, nil
}
