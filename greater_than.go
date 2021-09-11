package lep

type GreaterThan statement

var _ Expression = (*GreaterThan)(nil)

func NewGreaterThan(param *Param, value Value) *GreaterThan {
	return &GreaterThan{
		Param: param,
		Value: value,
	}
}

func (s GreaterThan) Equals(other Expression) bool {
	if expr, ok := other.(*GreaterThan); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s GreaterThan) String() string {
	return s.Param.String() + ">" + s.Value.String()
}

func parseGreaterThan(left, right interface{}) (*GreaterThan, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NewGreaterThan(param, value), nil
}

type GreaterThanEqual statement

var _ Expression = (*GreaterThanEqual)(nil)

func NewGreaterThanEqual(param *Param, value Value) *GreaterThanEqual {
	return &GreaterThanEqual{
		Param: param,
		Value: value,
	}
}

func (s GreaterThanEqual) Equals(other Expression) bool {
	if expr, ok := other.(*GreaterThanEqual); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s GreaterThanEqual) String() string {
	return s.Param.String() + ">=" + s.Value.String()
}

func parseGreaterThanEqual(left, right interface{}) (*GreaterThanEqual, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NewGreaterThanEqual(param, value), nil
}
