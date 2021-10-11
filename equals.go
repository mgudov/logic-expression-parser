package lep

type EqualsX struct {
	statement
}

var _ Expression = (*EqualsX)(nil)
var _ Statement = (*EqualsX)(nil)

func Equals(param *ParamX, value Value) *EqualsX {
	return &EqualsX{
		statement: statement{
			Param: param,
			Value: value,
		},
	}
}

func (e EqualsX) Equals(other Expression) bool {
	if expr, ok := other.(*EqualsX); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e EqualsX) String() string {
	return e.Param.String() + "=" + e.Value.String()
}

func parseEquals(left, right interface{}) (*EqualsX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return Equals(param, value), nil
}

type NotEqualsX struct {
	statement
}

var _ Expression = (*NotEqualsX)(nil)
var _ Statement = (*NotEqualsX)(nil)

func NotEquals(param *ParamX, value Value) *NotEqualsX {
	return &NotEqualsX{
		statement: statement{
			Param: param,
			Value: value,
		},
	}
}

func (e NotEqualsX) Equals(other Expression) bool {
	if expr, ok := other.(*NotEqualsX); ok {
		return e.Param.Equals(expr.Param) && e.Value.Equals(expr.Value)
	}
	return false
}

func (e NotEqualsX) String() string {
	return e.Param.String() + "!=" + e.Value.String()
}

func parseNotEquals(left, right interface{}) (*NotEqualsX, error) {
	param, value, err := parseStatement(left, right)
	if err != nil {
		return nil, err
	}
	return NotEquals(param, value), nil
}
