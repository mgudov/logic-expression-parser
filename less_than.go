package lep

type LessThan statement

var _ Expression = (*LessThan)(nil)

func NewLessThan(param *Param, value Value) *LessThan {
	return &LessThan{
		Param: param,
		Value: value,
	}
}

func (s LessThan) Equals(other Expression) bool {
	if expr, ok := other.(*LessThan); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s LessThan) String() string {
	return s.Param.String() + "<" + s.Value.String()
}

func parseLessThan(left, right interface{}) (*LessThan, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NewLessThan(param, value), nil
}

type LessThanEqual statement

var _ Expression = (*LessThanEqual)(nil)

func NewLessThanEqual(param *Param, value Value) *LessThanEqual {
	return &LessThanEqual{
		Param: param,
		Value: value,
	}
}

func (s LessThanEqual) Equals(other Expression) bool {
	if expr, ok := other.(*LessThanEqual); ok {
		return s.Param.Equals(expr.Param) && s.Value.Equals(expr.Value)
	}
	return false
}

func (s LessThanEqual) String() string {
	return s.Param.String() + "<=" + s.Value.String()
}

func parseLessThanEqual(left, right interface{}) (*LessThanEqual, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NewLessThanEqual(param, value), nil
}
