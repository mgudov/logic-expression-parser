package lep

type GreaterThanX statement

var _ Expression = (*GreaterThanX)(nil)

func GreaterThan(param *ParamX, value Value) *GreaterThanX {
	return &GreaterThanX{
		Param: param,
		Value: value,
	}
}

func (s GreaterThanX) Equals(other Expression) bool {
	if expr, ok := other.(*GreaterThanX); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s GreaterThanX) String() string {
	return s.Param.String() + ">" + s.Value.String()
}

func parseGreaterThan(left, right interface{}) (*GreaterThanX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return GreaterThan(param, value), nil
}

type GreaterThanEqualX statement

var _ Expression = (*GreaterThanEqualX)(nil)

func GreaterThanEqual(param *ParamX, value Value) *GreaterThanEqualX {
	return &GreaterThanEqualX{
		Param: param,
		Value: value,
	}
}

func (s GreaterThanEqualX) Equals(other Expression) bool {
	if expr, ok := other.(*GreaterThanEqualX); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s GreaterThanEqualX) String() string {
	return s.Param.String() + ">=" + s.Value.String()
}

func parseGreaterThanEqual(left, right interface{}) (*GreaterThanEqualX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return GreaterThanEqual(param, value), nil
}
