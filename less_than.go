package lep

type LessThanX statement

var _ Expression = (*LessThanX)(nil)

func LessThan(param *ParamX, value Value) *LessThanX {
	return &LessThanX{
		Param: param,
		Value: value,
	}
}

func (s LessThanX) Equals(other Expression) bool {
	if expr, ok := other.(*LessThanX); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s LessThanX) String() string {
	return s.Param.String() + "<" + s.Value.String()
}

func parseLessThan(left, right interface{}) (*LessThanX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return LessThan(param, value), nil
}

type LessThanEqualX statement

var _ Expression = (*LessThanEqualX)(nil)

func LessThanEqual(param *ParamX, value Value) *LessThanEqualX {
	return &LessThanEqualX{
		Param: param,
		Value: value,
	}
}

func (s LessThanEqualX) Equals(other Expression) bool {
	if expr, ok := other.(*LessThanEqualX); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s LessThanEqualX) String() string {
	return s.Param.String() + "<=" + s.Value.String()
}

func parseLessThanEqual(left, right interface{}) (*LessThanEqualX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return LessThanEqual(param, value), nil
}
